package middlewares

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/dtome123/auth-sdk/jwtutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthMode int

const (
	AuthNone AuthMode = iota
	AuthService
	AuthUser
)

type MethodRule struct {
	Pattern string
	Mode    AuthMode
}

type OauthInterceptor struct {
	audience         string
	verifier         jwtutils.JWTVerifier
	methodRules      []MethodRule
	replayChecker    jwtutils.ReplayChecker
	attachIdentityFn func(context.Context, jwtutils.Claims) context.Context
	tokenHeader      string
	tokenScheme      string
}

// === Functional Options ===

type OauthInterceptorOption func(*OauthInterceptor)

func WithMethodRules(rules []MethodRule) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.methodRules = rules
	}
}

func WithGRPCReplayChecker(ttl time.Duration) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.replayChecker = jwtutils.NewMemoryReplayChecker(ttl)
	}
}

func WithGRPCIdentityInjector(fn func(context.Context, jwtutils.Claims) context.Context) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.attachIdentityFn = fn
	}
}

func WithGRPCTokenHeader(header string) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.tokenHeader = strings.ToLower(header)
	}
}

func WithGRPCTokenScheme(scheme string) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.tokenScheme = strings.ToLower(scheme)
	}
}

// === Constructor ===

func NewOauthInterceptor(audience string, verifier jwtutils.JWTVerifier, opts ...OauthInterceptorOption) *OauthInterceptor {
	o := &OauthInterceptor{
		audience:    audience,
		verifier:    verifier,
		methodRules: []MethodRule{},
		tokenHeader: "authorization",
		tokenScheme: "bearer",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// === Unary Interceptor ===

func (o *OauthInterceptor) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		mode := o.matchMethod(info.FullMethod)

		if mode == AuthNone {
			return handler(ctx, req)
		}

		tokenStr, err := o.extractToken(ctx)
		if err != nil {
			return nil, err
		}

		claims, err := o.verifyTokenByMode(tokenStr, mode)
		if err != nil {
			return nil, err
		}

		if o.attachIdentityFn != nil {
			ctx = o.attachIdentityFn(ctx, claims)
		}

		return handler(ctx, req)
	}
}

// === Helpers ===

func (o *OauthInterceptor) matchMethod(method string) AuthMode {
	for _, rule := range o.methodRules {
		if matchPattern(method, rule.Pattern) {
			return rule.Mode
		}
	}
	return AuthNone
}

func (o *OauthInterceptor) verifyTokenByMode(token string, mode AuthMode) (jwtutils.Claims, error) {
	switch mode {
	case AuthService:
		return verifyServiceToken(o.verifier, token, o.audience, o.replayChecker)
	case AuthUser:
		return verifyUserToken(o.verifier, token, o.audience)
	default:
		return nil, errors.New("unsupported auth mode")
	}
}

func (o *OauthInterceptor) extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing metadata in context")
	}

	values := md.Get(o.tokenHeader)
	if len(values) == 0 {
		return "", errors.New("missing token header: " + o.tokenHeader)
	}

	raw := strings.TrimSpace(values[0])
	if o.tokenScheme == "" {
		if raw == "" {
			return "", errors.New("empty token")
		}
		return raw, nil
	}

	prefix := strings.ToLower(o.tokenScheme) + " "
	if !strings.HasPrefix(strings.ToLower(raw), prefix) {
		return "", errors.New("invalid token scheme, expected: " + o.tokenScheme)
	}

	token := strings.TrimSpace(raw[len(prefix):])
	if token == "" {
		return "", errors.New("empty token after scheme")
	}

	return token, nil
}

// === Pattern matcher with "*" suffix support ===

func matchPattern(s, pattern string) bool {
	if strings.HasSuffix(pattern, "*") {
		prefix := strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(s, prefix)
	}
	return s == pattern
}
