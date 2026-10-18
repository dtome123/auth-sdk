package middlewares

import (
	"context"

	"github.com/dtome123/auth-sdk/jwt"
	"github.com/dtome123/auth-sdk/session"
	"github.com/gin-gonic/gin"
)

type claimsCtxKey struct{}
type rawTokenCtxKey struct{}
type sessionCtxKey struct{}

var (
	claimsKey   = claimsCtxKey{}
	rawTokenKey = rawTokenCtxKey{}
	sessionKey  = sessionCtxKey{}
)

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

// ContextWithRawToken stores the raw token string inside a standard context.Context.
func ContextWithRawToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, rawTokenKey, token)
}

// RawTokenFromContext retrieves the raw token string from a standard context.Context.
func RawTokenFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	token, ok := ctx.Value(rawTokenKey).(string)
	return token, ok
}

// ContextWithSession stores a session entity pointer inside a standard context.Context.
func ContextWithSession[T any](ctx context.Context, sess *session.Session[T]) context.Context {
	return context.WithValue(ctx, sessionKey, sess)
}

// SessionFromContext retrieves a session entity pointer of type *session.Session[T] from a standard context.Context.
func SessionFromContext[T any](ctx context.Context) (*session.Session[T], bool) {
	if ctx == nil {
		return nil, false
	}
	val := ctx.Value(sessionKey)
	if val == nil {
		return nil, false
	}
	sess, ok := val.(*session.Session[T])
	return sess, ok
}

// GinClaimsKey is the key used to store Claims in *gin.Context.
const GinClaimsKey = "claims"

// GinRawTokenKey is the key used to store the raw token string in *gin.Context.
const GinRawTokenKey = "raw_token"

// GinSessionKey is the key used to store the session entity in *gin.Context.
const GinSessionKey = "session"

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

// SetGinRawToken stores the raw token string inside *gin.Context.
func SetGinRawToken(c *gin.Context, token string) {
	c.Set(GinRawTokenKey, token)
}

// GetGinRawToken retrieves the raw token string from *gin.Context.
func GetGinRawToken(c *gin.Context) (string, bool) {
	val, exists := c.Get(GinRawTokenKey)
	if !exists {
		return "", false
	}
	token, ok := val.(string)
	return token, ok
}

// SetGinSession stores a session entity pointer inside *gin.Context.
func SetGinSession[T any](c *gin.Context, sess *session.Session[T]) {
	c.Set(GinSessionKey, sess)
}

// GetGinSession retrieves a session entity pointer of type *session.Session[T] from *gin.Context.
func GetGinSession[T any](c *gin.Context) (*session.Session[T], bool) {
	val, exists := c.Get(GinSessionKey)
	if !exists || val == nil {
		return nil, false
	}
	sess, ok := val.(*session.Session[T])
	return sess, ok
}
