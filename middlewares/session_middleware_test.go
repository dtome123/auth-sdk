package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dtome123/auth-sdk/jwt"
	"github.com/dtome123/auth-sdk/session"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testUserPayload struct {
	Role string `json:"role"`
}

type testSessionRepo struct {
	sessions map[string]*session.Session[testUserPayload]
}

func (m *testSessionRepo) GetSession(ctx context.Context, userID, deviceID string) (*session.Session[testUserPayload], error) {
	key := userID + ":" + deviceID
	sess, ok := m.sessions[key]
	if !ok {
		return nil, nil
	}
	return sess, nil
}

func (m *testSessionRepo) UpsertSession(ctx context.Context, s session.Session[testUserPayload]) error {
	key := s.UserID + ":" + s.DeviceID
	m.sessions[key] = &s
	return nil
}

func (m *testSessionRepo) DeleteSession(ctx context.Context, userID, deviceID string) error {
	delete(m.sessions, userID+":"+deviceID)
	return nil
}

func TestCheckSessionGin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &testSessionRepo{
		sessions: map[string]*session.Session[testUserPayload]{
			"u1:d1": {
				BaseSession: session.BaseSession{UserID: "u1", DeviceID: "d1", JTI: "jti-1"},
				Data:        testUserPayload{Role: "admin"},
			},
		},
	}

	t.Run("Success with defaults", func(t *testing.T) {
		r := gin.New()
		r.Use(func(c *gin.Context) {
			SetGinClaims(c, jwt.Claims{
				"sub":       jwt.Value{Value: "u1"},
				"device_id": jwt.Value{Value: "d1"},
				"jti":       jwt.Value{Value: "jti-1"},
			})
			c.Next()
		})
		r.Use(CheckSession(SessionConfig[testUserPayload]{Repo: repo}))
		r.GET("/test", func(c *gin.Context) {
			sess, ok := GetGinSession[testUserPayload](c)
			if !ok || sess == nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			c.String(http.StatusOK, sess.Data.Role)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK || w.Body.String() != "admin" {
			t.Fatalf("expected 200 admin, got %d %s", w.Code, w.Body.String())
		}
	})

	t.Run("JTI Mismatch", func(t *testing.T) {
		r := gin.New()
		r.Use(func(c *gin.Context) {
			SetGinClaims(c, jwt.Claims{
				"sub":       jwt.Value{Value: "u1"},
				"device_id": jwt.Value{Value: "d1"},
				"jti":       jwt.Value{Value: "jti-other"},
			})
			c.Next()
		})
		r.Use(CheckSession(SessionConfig[testUserPayload]{Repo: repo}))
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", w.Code)
		}
	})

	t.Run("Success with custom claims", func(t *testing.T) {
		r := gin.New()
		r.Use(func(c *gin.Context) {
			SetGinClaims(c, jwt.Claims{
				"custom_uid": jwt.Value{Value: "u1"},
				"custom_did": jwt.Value{Value: "d1"},
				"custom_jti": jwt.Value{Value: "jti-1"},
			})
			c.Next()
		})
		r.Use(CheckSession(SessionConfig[testUserPayload]{
			Repo:          repo,
			UserIDClaim:   "custom_uid",
			DeviceIDClaim: "custom_did",
			JTIClaim:      "custom_jti",
		}))
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Session Not Found", func(t *testing.T) {
		r := gin.New()
		r.Use(func(c *gin.Context) {
			SetGinClaims(c, jwt.Claims{
				"sub":       jwt.Value{Value: "u_none"},
				"device_id": jwt.Value{Value: "d_none"},
				"jti":       jwt.Value{Value: "jti-none"},
			})
			c.Next()
		})
		r.Use(CheckSession(SessionConfig[testUserPayload]{Repo: repo}))
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", w.Code)
		}
	})
}

func TestCheckSessionGRPC(t *testing.T) {
	repo := &testSessionRepo{
		sessions: map[string]*session.Session[testUserPayload]{
			"u1:d1": {
				BaseSession: session.BaseSession{UserID: "u1", DeviceID: "d1", JTI: "jti-grpc"},
				Data:        testUserPayload{Role: "member"},
			},
		},
	}

	handler := func(ctx context.Context, req any) (any, error) {
		sess, ok := SessionFromContext[testUserPayload](ctx)
		if !ok || sess == nil {
			return nil, status.Error(codes.Internal, "no session")
		}
		return sess.Data.Role, nil
	}
	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Get"}

	t.Run("Success", func(t *testing.T) {
		ctx := ContextWithClaims(context.Background(), jwt.Claims{
			"sub":       jwt.Value{Value: "u1"},
			"device_id": jwt.Value{Value: "d1"},
			"jti":       jwt.Value{Value: "jti-grpc"},
		})

		interceptor := CheckSessionInterceptor(SessionConfig[testUserPayload]{Repo: repo})
		resp, err := interceptor(ctx, "req", info, handler)
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}
		if resp != "member" {
			t.Fatalf("expected member, got %v", resp)
		}
	})

	t.Run("JTI Mismatch", func(t *testing.T) {
		ctx := ContextWithClaims(context.Background(), jwt.Claims{
			"sub":       jwt.Value{Value: "u1"},
			"device_id": jwt.Value{Value: "d1"},
			"jti":       jwt.Value{Value: "jti-mismatch"},
		})

		interceptor := CheckSessionInterceptor(SessionConfig[testUserPayload]{Repo: repo})
		_, err := interceptor(ctx, "req", info, handler)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		st, _ := status.FromError(err)
		if st.Code() != codes.Unauthenticated {
			t.Fatalf("expected Unauthenticated, got %v", st.Code())
		}
	})

	t.Run("Session Not Found", func(t *testing.T) {
		ctx := ContextWithClaims(context.Background(), jwt.Claims{
			"sub":       jwt.Value{Value: "u_unknown"},
			"device_id": jwt.Value{Value: "d1"},
			"jti":       jwt.Value{Value: "jti-unknown"},
		})

		interceptor := CheckSessionInterceptor(SessionConfig[testUserPayload]{Repo: repo})
		_, err := interceptor(ctx, "req", info, handler)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		st, _ := status.FromError(err)
		if st.Code() != codes.Unauthenticated {
			t.Fatalf("expected Unauthenticated, got %v", st.Code())
		}
	})
}
