package middlewares

import (
	"errors"
	"strings"

	"github.com/dtome123/auth-sdk/assertion"
	"github.com/dtome123/auth-sdk/jwt"
	"github.com/gin-gonic/gin"
)

// === Middleware struct ===

type OauthMiddleware struct {
	audience           string
	verifier           jwt.JWTVerifier
	replayChecker      assertion.ReplayChecker
	attachIdentityFn   func(c *gin.Context, claims jwt.Claims)
	tokenHeader        string // default: "Authorization"
	tokenScheme        string // default: "Bearer"
	basicAuthAccounts  map[string]string
	basicAuthValidator func(username, password string) (jwt.Claims, bool)
	cookieName         string // default: "access_token"
}

// === Functional Options ===

type OauthMiddlewareOption func(*OauthMiddleware)

func WithReplayChecker(checker assertion.ReplayChecker) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.replayChecker = checker
	}
}

func WithIdentityInjector(fn func(*gin.Context, jwt.Claims)) OauthMiddlewareOption {
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

func WithBasicAuthAccounts(accounts map[string]string) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.basicAuthAccounts = accounts
	}
}

func WithBasicAuthValidator(fn func(username, password string) (jwt.Claims, bool)) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.basicAuthValidator = fn
	}
}

func WithCookieName(name string) OauthMiddlewareOption {
	return func(o *OauthMiddleware) {
		o.cookieName = name
	}
}

// === Constructor ===

func NewOauthMiddleware(
	audience string,
	verifier jwt.JWTVerifier,
	opts ...OauthMiddlewareOption,
) *OauthMiddleware {
	o := &OauthMiddleware{
		audience:    audience,
		verifier:    verifier,
		tokenHeader: "Authorization",
		tokenScheme: "Bearer",
		cookieName:  "access_token",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// === Core verification middleware logic ===

func (o *OauthMiddleware) VerifyToken(opts ...jwt.ValidateOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims jwt.Claims

		scheme, payload, extractErr := o.extractTokenAndScheme(c)
		if extractErr == nil {
			switch scheme {
			case "basic":
				if c.Request == nil {
					c.AbortWithStatusJSON(401, gin.H{"error": "missing request context"})
					return
				}
				user, pass, ok := c.Request.BasicAuth()
				if !ok {
					c.AbortWithStatusJSON(401, gin.H{"error": "invalid basic auth header"})
					return
				}
				cClaims, valid := o.validateBasicAuth(user, pass)
				if !valid {
					c.AbortWithStatusJSON(401, gin.H{"error": "invalid basic auth credentials"})
					return
				}
				claims = cClaims

			default:
				if payload == "" {
					c.AbortWithStatusJSON(401, gin.H{"error": "empty token"})
					return
				}
				cClaims, err := verifyTokenByMode(o.verifier, payload, o.audience, o.replayChecker, opts)
				if err != nil {
					c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
					return
				}
				claims = cClaims
			}
		} else {
			// Fallback to Cookie Auth if tokenHeader is missing
			if o.cookieName != "" {
				cookieToken, cookieErr := c.Cookie(o.cookieName)
				if cookieErr == nil && cookieToken != "" {
					cClaims, err := verifyTokenByMode(o.verifier, cookieToken, o.audience, o.replayChecker, opts)
					if err == nil {
						claims = cClaims
					} else {
						c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
						return
					}
				} else {
					c.AbortWithStatusJSON(401, gin.H{"error": extractErr.Error()})
					return
				}
			} else {
				c.AbortWithStatusJSON(401, gin.H{"error": extractErr.Error()})
				return
			}
		}

		if o.attachIdentityFn != nil && claims != nil {
			o.attachIdentityFn(c, claims)
		}

		c.Next()
	}
}

func (o *OauthMiddleware) validateBasicAuth(user, pass string) (jwt.Claims, bool) {
	if o.basicAuthValidator != nil {
		return o.basicAuthValidator(user, pass)
	}
	if len(o.basicAuthAccounts) > 0 {
		expectedPass, exists := o.basicAuthAccounts[user]
		if exists && expectedPass == pass {
			claims := jwt.Claims{
				"sub": jwt.Value{Value: user},
			}
			return claims, true
		}
	}
	return nil, false
}

// === Header & Scheme extractor ===

func (o *OauthMiddleware) extractTokenAndScheme(c *gin.Context) (scheme string, payload string, err error) {
	rawHeader := c.GetHeader(o.tokenHeader)
	if rawHeader == "" {
		return "", "", errors.New("missing token header: " + o.tokenHeader)
	}

	rawTrimmed := strings.TrimSpace(rawHeader)
	parts := strings.SplitN(rawTrimmed, " ", 2)
	if len(parts) == 1 {
		return strings.ToLower(o.tokenScheme), parts[0], nil
	}

	return strings.ToLower(parts[0]), strings.TrimSpace(parts[1]), nil
}
