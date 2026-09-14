package middlewares

import (
	"context"

	"github.com/dtome123/auth-sdk/jwt"
	"github.com/gin-gonic/gin"
)

type claimsCtxKey struct{}

var claimsKey = claimsCtxKey{}

// ContextWithClaims stores jwt.Claims inside a standard context.Context.
func ContextWithClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// ClaimsFromContext retrieves jwt.Claims from a standard context.Context.
func ClaimsFromContext(ctx context.Context) (jwt.Claims, bool) {
	if ctx == nil {
		return nil, false
	}
	claims, ok := ctx.Value(claimsKey).(jwt.Claims)
	return claims, ok
}

// GinClaimsKey is the key used to store Claims in *gin.Context.
const GinClaimsKey = "claims"

// SetGinClaims stores jwt.Claims inside *gin.Context.
func SetGinClaims(c *gin.Context, claims jwt.Claims) {
	c.Set(GinClaimsKey, claims)
}

// GetGinClaims retrieves jwt.Claims from *gin.Context.
func GetGinClaims(c *gin.Context) (jwt.Claims, bool) {
	val, exists := c.Get(GinClaimsKey)
	if !exists {
		return nil, false
	}
	claims, ok := val.(jwt.Claims)
	return claims, ok
}
