package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/dtome123/auth-sdk/assertion"
	"github.com/dtome123/auth-sdk/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ======================
// Auth Interceptor (Authentication / AuthN)
// ======================

// AuthInterceptor handles extracting and verifying tokens for incoming gRPC calls.
type AuthInterceptor struct {
	audience         string
	verifier         jwt.Verifier
	replayChecker    assertion.ReplayChecker
	attachIdentityFn func(context.Context, jwt.Claims) context.Context
	tokenHeader      string
	tokenScheme      string
}

// Backward compatibility alias
type OauthInterceptor = AuthInterceptor

// Functional Options
type AuthInterceptorOption func(*AuthInterceptor)
type OauthInterceptorOption = AuthInterceptorOption

func WithGRPCReplayChecker(checker assertion.ReplayChecker) AuthInterceptorOption {
	return func(o *AuthInterceptor) {
		o.replayChecker = checker
	}
}

func WithGRPCIdentityInjector(fn func(context.Context, jwt.Claims) context.Context) AuthInterceptorOption {
	return func(o *AuthInterceptor) {
		o.attachIdentityFn = fn
	}
}

func WithGRPCTokenHeader(header string) AuthInterceptorOption {
	return func(o *AuthInterceptor) {
		o.tokenHeader = strings.ToLower(header)
	}
}

func WithGRPCTokenScheme(scheme string) AuthInterceptorOption {
	return func(o *AuthInterceptor) {
		o.tokenScheme = strings.ToLower(scheme)
	}
}

// NewAuthInterceptor creates a new gRPC authentication interceptor.
func NewAuthInterceptor(audience string, verifier jwt.Verifier, opts ...AuthInterceptorOption) *AuthInterceptor {
	o := &AuthInterceptor{
		audience:    audience,
		verifier:    verifier,
		tokenHeader: "authorization",
		tokenScheme: "bearer",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// NewOauthInterceptor is an alias for NewAuthInterceptor for backward compatibility.
func NewOauthInterceptor(audience string, verifier jwt.Verifier, opts ...AuthInterceptorOption) *AuthInterceptor {
	return NewAuthInterceptor(audience, verifier, opts...)
}

// UnaryServerInterceptor returns the pure Authentication (AuthN) interceptor.
// It verifies the token and injects Claims into context.
// Returns codes.Unauthenticated if token is invalid or missing.
func (a *AuthInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.VerifyInterceptor()
}

// VerifyInterceptor verifies the authentication of the incoming gRPC request.
// Preserved for backward compatibility.
func (a *AuthInterceptor) VerifyInterceptor(opts ...jwt.ValidateOption) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata in context")
		}

		_, payload, err := a.extractTokenAndScheme(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		if payload == "" {
			return nil, status.Error(codes.Unauthenticated, "empty token")
		}

		claims, err := verifyToken(a.verifier, payload, a.audience, a.replayChecker, opts)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Inject claims into context
		ctx = ContextWithClaims(ctx, claims)
		ctx = ContextWithRawToken(ctx, payload)
		if a.attachIdentityFn != nil {
			ctx = a.attachIdentityFn(ctx, claims)
		}

		return handler(ctx, req)
	}
}

// === Helpers ===

func (a *AuthInterceptor) extractTokenAndScheme(md metadata.MD) (scheme string, payload string, err error) {
	values := md.Get(a.tokenHeader)
	if len(values) == 0 {
		return "", "", errors.New("missing token header: " + a.tokenHeader)
	}

	raw := strings.TrimSpace(values[0])
	parts := strings.SplitN(raw, " ", 2)
	if len(parts) == 1 {
		return strings.ToLower(a.tokenScheme), parts[0], nil
	}

	return strings.ToLower(parts[0]), strings.TrimSpace(parts[1]), nil
}

// matchPattern supports exact match or "*" suffix matching.
func matchPattern(s, pattern string) bool {
	if prefix, ok := strings.CutSuffix(pattern, "*"); ok {
		return strings.HasPrefix(s, prefix)
	}
	return s == pattern
}
