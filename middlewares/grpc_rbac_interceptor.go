package middlewares

import (
	"context"
	"slices"

	"github.com/dtome123/auth-sdk/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ======================
// RBAC & Authorization Interceptors (AuthZ)
// ======================

// AuthorizationRule defines the authorization criteria for gRPC methods matching Pattern.
type AuthorizationRule struct {
	Pattern          string // Method pattern, e.g. "/auth.v1.TokenService/*" or exact name
	RequireRoles     []string
	RequireScopes    []string
	RequirePerms     []string
	CustomAuthorizer func(ctx context.Context, claims jwt.Claims) (bool, error)
	Skip             bool
}

// AuthorizationInterceptor enforces method-level authorization for gRPC calls.
type AuthorizationInterceptor struct {
	rules []AuthorizationRule
}

// NewAuthorizationInterceptor creates a new gRPC authorization interceptor.
func NewAuthorizationInterceptor(rules ...AuthorizationRule) *AuthorizationInterceptor {
	return &AuthorizationInterceptor{
		rules: rules,
	}
}

// UnaryServerInterceptor returns the authorization unary server interceptor.
// Returns codes.Unauthenticated if claims are missing from context,
// and codes.PermissionDenied if the claims do not satisfy the required rules.
func (a *AuthorizationInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		rule := a.matchRule(info.FullMethod)
		if rule == nil || rule.Skip {
			return handler(ctx, req)
		}

		claims, ok := ClaimsFromContext(ctx)
		if !ok || claims == nil {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated: missing claims in context")
		}

		// 1. Role verification
		if len(rule.RequireRoles) > 0 {
			userRoles := extractRolesFromClaims(claims)
			matched := false
			for _, reqRole := range rule.RequireRoles {
				if slices.Contains(userRoles, reqRole) {
					matched = true
					break
				}
			}
			if !matched {
				return nil, status.Error(codes.PermissionDenied, "permission denied: missing required role")
			}
		}

		// 2. Scope verification
		if len(rule.RequireScopes) > 0 {
			tokenScopes := extractScopesFromClaims(claims)
			for _, reqScope := range rule.RequireScopes {
				if !slices.Contains(tokenScopes, reqScope) {
					return nil, status.Errorf(codes.PermissionDenied, "permission denied: missing required scope: %s", reqScope)
				}
			}
		}

		// 3. Permission verification
		if len(rule.RequirePerms) > 0 {
			userPerms := extractPermissionsFromClaims(claims)
			for _, reqPerm := range rule.RequirePerms {
				if !slices.Contains(userPerms, reqPerm) {
					return nil, status.Errorf(codes.PermissionDenied, "permission denied: missing required permission: %s", reqPerm)
				}
			}
		}

		// 4. Custom authorizer function
		if rule.CustomAuthorizer != nil {
			allowed, err := rule.CustomAuthorizer(ctx, claims)
			if err != nil {
				return nil, status.Errorf(codes.PermissionDenied, "permission denied: %v", err)
			}
			if !allowed {
				return nil, status.Error(codes.PermissionDenied, "permission denied: access denied")
			}
		}

		return handler(ctx, req)
	}
}

func (a *AuthorizationInterceptor) matchRule(method string) *AuthorizationRule {
	for _, rule := range a.rules {
		if matchPattern(method, rule.Pattern) {
			return &rule
		}
	}
	return nil
}
