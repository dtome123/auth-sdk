package middlewares

import (
	"errors"
	"strings"
	"time"

	"github.com/dtome123/auth-sdk/jwtutils"
	"github.com/gin-gonic/gin"
)

// === Middleware struct ===

type OauthMiddleware struct {
	audience         string
	verifier         jwtutils.JWTVerifier
	replayChecker    jwtutils.ReplayChecker
	attachIdentityFn func(c *gin.Context, claims jwtutils.Claims)
	tokenHeader      string // default: "Authorization"
	tokenScheme      string // default: "Bearer"
}

// === Functional Options ===

type OauthMiddlewareOption func(*OauthMiddleware)

func WithReplayChecker(ttl time.Duration) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.replayChecker = jwtutils.NewMemoryReplayChecker(ttl)
	}
}

func WithIdentityInjector(fn func(*gin.Context, jwtutils.Claims)) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.attachIdentityFn = fn
	}
}

func WithTokenHeader(header string) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.tokenHeader = header
	}
}

func WithTokenScheme(scheme string) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.tokenScheme = scheme
	}
}

// === Constructor ===

func NewOauthMiddleware(
	audience string,
	verifier jwtutils.JWTVerifier,
	opts ...OauthMiddlewareOption,
) *OauthMiddleware {
	o := &OauthMiddleware{
		audience:    audience,
		verifier:    verifier,
		tokenHeader: "Authorization",
		tokenScheme: "Bearer",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// === Middleware for service token ===

func (o *OauthMiddleware) ServiceMiddleware() gin.HandlerFunc {
	return o.middlewareWithVerifier(func(token string) (jwtutils.Claims, error) {
		return verifyServiceToken(o.verifier, token, o.audience, o.replayChecker)
	})
}

// === Middleware for user token ===

func (o *OauthMiddleware) UserMiddleware() gin.HandlerFunc {
	return o.middlewareWithVerifier(func(token string) (jwtutils.Claims, error) {
		return verifyUserToken(o.verifier, token, o.audience)
	})
}

// === Core verification middleware logic ===

func (o *OauthMiddleware) middlewareWithVerifier(
	verifyFn func(tokenStr string) (jwtutils.Claims, error),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := o.extractToken(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		claims, err := verifyFn(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if o.attachIdentityFn != nil {
			o.attachIdentityFn(c, claims)
		}

		c.Next()
	}
}

// === Token extractor (header + scheme insensitive) ===

func (o *OauthMiddleware) extractToken(c *gin.Context) (string, error) {
	rawHeader := c.GetHeader(o.tokenHeader)
	if rawHeader == "" {
		return "", errors.New("missing token header: " + o.tokenHeader)
	}

	if o.tokenScheme == "" {
		return strings.TrimSpace(rawHeader), nil
	}

	prefix := o.tokenScheme + " "
	if !strings.HasPrefix(strings.ToLower(rawHeader), strings.ToLower(prefix)) {
		return "", errors.New("invalid token scheme, expected: " + o.tokenScheme)
	}

	token := strings.TrimSpace(rawHeader[len(prefix):])
	if token == "" {
		return "", errors.New("empty token after scheme")
	}
	return token, nil
}
