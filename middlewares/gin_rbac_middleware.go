package middlewares

import (
	"net/http"
	"slices"
	"strings"

	"github.com/dtome123/auth-sdk/jwt"
	"github.com/gin-gonic/gin"
)

// ======================
// RBAC & Authorization Middlewares (AuthZ)
// ======================

// AuthorizerFunc is a custom authorization check function.
type AuthorizerFunc func(c *gin.Context, claims jwt.Claims) (bool, error)

// Authorize checks whether the authenticated user passes the provided authorizer function.
// Returns HTTP 401 if unauthenticated, and HTTP 403 if permission is denied.
func Authorize(fn AuthorizerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetGinClaims(c)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated: missing claims"})
			return
		}

		allowed, err := fn(c, claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: " + err.Error()})
			return
		}
		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: access denied"})
			return
		}

		c.Next()
	}
}

// RequireRoles checks if the authenticated user has at least one of the required roles.
// Looks in claims: "roles" (array/string) or "role" (string).
func RequireRoles(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetGinClaims(c)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated: missing claims"})
			return
		}

		userRoles := extractRolesFromClaims(claims)
		matched := false
		for _, required := range requiredRoles {
			if slices.Contains(userRoles, required) {
				matched = true
				break
			}
		}

		if !matched {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: missing required role"})
			return
		}

		c.Next()
	}
}

// RequireScopes checks if the authenticated user has all of the required scopes.
// Looks in claims: "scope" (space-separated string or slice) or "scopes" or "scp".
func RequireScopes(requiredScopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetGinClaims(c)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated: missing claims"})
			return
		}

		tokenScopes := extractScopesFromClaims(claims)
		for _, reqScope := range requiredScopes {
			if !slices.Contains(tokenScopes, reqScope) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: missing required scope: " + reqScope})
				return
			}
		}

		c.Next()
	}
}

// RequirePermissions checks if the authenticated user has all of the required permissions.
// Looks in claims: "permissions" (slice/string) or "perms".
func RequirePermissions(requiredPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetGinClaims(c)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated: missing claims"})
			return
		}

		userPerms := extractPermissionsFromClaims(claims)
		for _, reqPerm := range requiredPermissions {
			if !slices.Contains(userPerms, reqPerm) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: missing required permission: " + reqPerm})
				return
			}
		}

		c.Next()
	}
}

// === Helper claim extractors ===

func extractRolesFromClaims(claims jwt.Claims) []string {
	if rolesVal := claims.Get("roles"); !rolesVal.IsNil() {
		return rolesVal.AsStringSlice()
	}
	if roleVal := claims.Get("role"); !roleVal.IsNil() && roleVal.AsString() != "" {
		return []string{roleVal.AsString()}
	}
	return nil
}

func extractScopesFromClaims(claims jwt.Claims) []string {
	if scopesVal := claims.Get("scopes"); !scopesVal.IsNil() {
		return scopesVal.AsStringSlice()
	}
	if scpVal := claims.Get("scp"); !scpVal.IsNil() {
		return scpVal.AsStringSlice()
	}
	if scopeVal := claims.Get("scope"); !scopeVal.IsNil() {
		// Scopes can be space-delimited string (RFC 6749) or slice
		rawStr := scopeVal.AsString()
		if rawStr != "" {
			return strings.Fields(rawStr)
		}
		return scopeVal.AsStringSlice()
	}
	return nil
}

func extractPermissionsFromClaims(claims jwt.Claims) []string {
	if permsVal := claims.Get("permissions"); !permsVal.IsNil() {
		return permsVal.AsStringSlice()
	}
	if permsVal := claims.Get("perms"); !permsVal.IsNil() {
		return permsVal.AsStringSlice()
	}
	return nil
}
