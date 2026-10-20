package middlewares

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckSessionInterceptor returns a gRPC unary interceptor that validates an active session for the authenticated user.
func CheckSessionInterceptor[T any](cfg SessionConfig[T]) grpc.UnaryServerInterceptor {

	userClaim := "sub"
	jtiClaim := "jti"

	deviceClaim := cfg.DeviceIDClaim
	if deviceClaim == "" {
		deviceClaim = "device_id"
	}

	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		claims, ok := ClaimsFromContext(ctx)
		if !ok || claims == nil {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated: missing claims in context")
		}

		userID := claims.Get(userClaim).AsString()
		deviceID := claims.Get(deviceClaim).AsString()
		jti := claims.Get(jtiClaim).AsString()

		if deviceID == "" {
			deviceID = "*"
		}

		if userID == "" || jti == "" {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated: missing user_id or jti in claims")
		}

		sess, err := cfg.Repo.GetSession(ctx, userID, deviceID)
		if err != nil || sess == nil {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated: session not found or expired")
		}

		if sess.JTI != jti {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated: session revoked or replaced")
		}

		ctx = ContextWithSession(ctx, sess)
		return handler(ctx, req)
	}
}
