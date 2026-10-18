package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dtome123/auth-sdk/assertion"
	"github.com/dtome123/auth-sdk/jwt"
	"github.com/gin-gonic/gin"
)

// ======================
// Auth Middleware (Authentication / AuthN)
// ======================

// AuthMiddleware handles extracting and verifying tokens for Gin HTTP requests.
type AuthMiddleware struct {
	audience           string
	verifier           jwt.Verifier
	replayChecker      assertion.ReplayChecker
	attachIdentityFn   func(c *gin.Context, claims jwt.Claims)
	tokenHeader        string // default: "Authorization"
	tokenScheme        string // default: "Bearer"
	basicAuthAccounts  map[string]string
	basicAuthValidator func(username, password string) (jwt.Claims, bool)
	cookieName         string // default: "access_token"
}

// Backward compatibility alias
type OauthMiddleware = AuthMiddleware

// Functional Option
type AuthMiddlewareOption func(*AuthMiddleware)
type OauthMiddlewareOption = AuthMiddlewareOption

func WithReplayChecker(checker assertion.ReplayChecker) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.replayChecker = checker
	}
}

func WithIdentityInjector(fn func(*gin.Context, jwt.Claims)) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.attachIdentityFn = fn
	}
}

func WithTokenHeader(header string) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.tokenHeader = header
	}
}

func WithTokenScheme(scheme string) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.tokenScheme = scheme
	}
}

func WithBasicAuthAccounts(accounts map[string]string) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.basicAuthAccounts = accounts
	}
}

func WithBasicAuthValidator(fn func(username, password string) (jwt.Claims, bool)) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.basicAuthValidator = fn
	}
}

func WithCookieName(name string) AuthMiddlewareOption {
	return func(o *AuthMiddleware) {
		o.cookieName = name
	}
}

// NewAuthMiddleware creates a new authentication middleware.
func NewAuthMiddleware(
	audience string,
	verifier jwt.Verifier,
	opts ...AuthMiddlewareOption,
) *AuthMiddleware {
	m := &AuthMiddleware{
		audience:    audience,
		verifier:    verifier,
		tokenHeader: "Authorization",
		tokenScheme: "Bearer",
		cookieName:  "access_token",
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// NewOauthMiddleware is an alias for NewAuthMiddleware for backward compatibility.
func NewOauthMiddleware(
	audience string,
	verifier jwt.Verifier,
	opts ...AuthMiddlewareOption,
) *AuthMiddleware {
	return NewAuthMiddleware(audience, verifier, opts...)
}

// Authenticate is the pure Authentication (AuthN) middleware.
// It verifies the identity and returns HTTP 401 Unauthorized if invalid or missing.
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := m.extractAndVerifyClaims(c, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Inject claims into context
		SetGinClaims(c, claims)
		if m.attachIdentityFn != nil && claims != nil {
			m.attachIdentityFn(c, claims)
		}

		c.Next()
	}
}

// VerifyToken maintains backward compatibility with previous API.
// It verifies token and optionally executes claims validation options.
func (m *AuthMiddleware) VerifyToken(opts ...jwt.ValidateOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := m.extractAndVerifyClaims(c, opts)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		SetGinClaims(c, claims)
		if m.attachIdentityFn != nil && claims != nil {
			m.attachIdentityFn(c, claims)
		}

		c.Next()
	}
}

// Helper to extract and verify claims
func (m *AuthMiddleware) extractAndVerifyClaims(c *gin.Context, validateOpts []jwt.ValidateOption) (jwt.Claims, error) {
	scheme, payload, extractErr := m.extractTokenAndScheme(c)
	if extractErr == nil {
		switch scheme {
		case "basic":
			if c.Request == nil {
				return nil, errors.New("missing request context")
			}
			user, pass, ok := c.Request.BasicAuth()
			if !ok {
				return nil, errors.New("invalid basic auth header")
			}
			cClaims, valid := m.validateBasicAuth(user, pass)
			if !valid {
				return nil, errors.New("invalid basic auth credentials")
			}
			return cClaims, nil

		default:
			if payload == "" {
				return nil, errors.New("empty token")
			}
			SetGinRawToken(c, payload)
			return verifyToken(m.verifier, payload, m.audience, m.replayChecker, validateOpts)
		}
	}

	// Fallback to Cookie Auth if tokenHeader is missing
	if m.cookieName != "" {
		cookieToken, cookieErr := c.Cookie(m.cookieName)
		if cookieErr == nil && cookieToken != "" {
			SetGinRawToken(c, cookieToken)
			return verifyToken(m.verifier, cookieToken, m.audience, m.replayChecker, validateOpts)
		}
	}

	return nil, extractErr
}

func (m *AuthMiddleware) validateBasicAuth(user, pass string) (jwt.Claims, bool) {
	if m.basicAuthValidator != nil {
		return m.basicAuthValidator(user, pass)
	}
	if len(m.basicAuthAccounts) > 0 {
		expectedPass, exists := m.basicAuthAccounts[user]
		if exists && expectedPass == pass {
			claims := jwt.Claims{
				"sub": jwt.Value{Value: user},
			}
			return claims, true
		}
	}
	return nil, false
}

func (m *AuthMiddleware) extractTokenAndScheme(c *gin.Context) (scheme string, payload string, err error) {
	rawHeader := c.GetHeader(m.tokenHeader)
	if rawHeader == "" {
		return "", "", errors.New("missing token header: " + m.tokenHeader)
	}

	rawTrimmed := strings.TrimSpace(rawHeader)
	parts := strings.SplitN(rawTrimmed, " ", 2)
	if len(parts) == 1 {
		return strings.ToLower(m.tokenScheme), parts[0], nil
	}

	return strings.ToLower(parts[0]), strings.TrimSpace(parts[1]), nil
}
