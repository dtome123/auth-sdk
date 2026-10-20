package middlewares

import (
	"context"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PermissionExtractor trích xuất danh sách permissions từ context.
type PermissionExtractor func(ctx context.Context) []string

// ResourceRBACConfig cấu hình kiểm tra quyền theo Resource-Action.
type ResourceRBACConfig struct {
	// Resource là định danh tài nguyên (ví dụ: "tenant:staff")
	Resource string

	// MethodActions map tên RPC method sang action tương ứng (ví dụ: "ListStaff": "read")
	MethodActions map[string]string

	// Extractor là hàm lấy permissions của user từ context
	Extractor PermissionExtractor
}

// RequireResourcePermission trả về unary interceptor kiểm tra quyền theo format <resource>:<action>.
func RequireResourcePermission(cfg ResourceRBACConfig) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := info.FullMethod
		if idx := strings.LastIndex(info.FullMethod, "/"); idx != -1 {
			methodName = info.FullMethod[idx+1:]
		}

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
