package middlewares

import (
	"context"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PermissionExtractor extracts user permissions from context.
type PermissionExtractor func(ctx context.Context) []string

// ResourceRBACConfig configures permission enforcement based on Resource-Action.
type ResourceRBACConfig struct {
	// Resource is the resource identifier (e.g. "tenant:staff")
	Resource string

	// MethodActions maps RPC method name to its required action (e.g. "ListStaff": "read")
	MethodActions map[string]string

	// Extractor extracts user permissions from context
	Extractor PermissionExtractor
}

// RequireResourcePermission returns a unary interceptor that enforces RBAC based on <resource>:<action>.
func RequireResourcePermission(cfg ResourceRBACConfig) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := info.FullMethod
		if idx := strings.LastIndex(info.FullMethod, "/"); idx != -1 {
			methodName = info.FullMethod[idx+1:]
		}

		// Deny by default: reject methods that have no configured action
		action, exists := cfg.MethodActions[methodName]
		if !exists || action == "" {
			return nil, status.Errorf(codes.PermissionDenied, "method %s has no permission configured", methodName)
		}

		requiredPermission := cfg.Resource + ":" + action
		adminPermission := cfg.Resource + ":*"

		var userPerms []string
		if cfg.Extractor != nil {
			userPerms = cfg.Extractor(ctx)
		}

		if slices.Contains(userPerms, requiredPermission) || slices.Contains(userPerms, adminPermission) {
			return handler(ctx, req)
		}

		return nil, status.Errorf(codes.PermissionDenied, "permission denied: require %s", requiredPermission)
	}
}
