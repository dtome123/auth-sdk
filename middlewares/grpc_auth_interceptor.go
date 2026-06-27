package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/dtome123/auth-sdk/assertion"
	"github.com/dtome123/auth-sdk/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MethodRule struct {
	Pattern         string
	Skip            bool
	ValidateOptions []jwt.ValidateOption
}

type OauthInterceptor struct {
	audience         string
	verifier         jwt.JWTVerifier
	methodRules      []MethodRule
	replayChecker    assertion.ReplayChecker
	attachIdentityFn func(context.Context, jwt.Claims) context.Context
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

func WithGRPCReplayChecker(checker assertion.ReplayChecker) OauthInterceptorOption {
	return func(o *OauthInterceptor) {
		o.replayChecker = checker
	}
}

func WithGRPCIdentityInjector(fn func(context.Context, jwt.Claims) context.Context) OauthInterceptorOption {
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

func NewOauthInterceptor(audience string, verifier jwt.JWTVerifier, opts ...OauthInterceptorOption) *OauthInterceptor {
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

func (o *OauthInterceptor) VerifyInterceptor(opts ...jwt.ValidateOption) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		rule := o.matchMethod(info.FullMethod)
		if rule != nil && rule.Skip {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata in context")
		}

		_, payload, err := o.extractTokenAndScheme(md)
		if err != nil {
			return nil, err
		}

		if payload == "" {
			return nil, errors.New("empty token")
		}

		combinedOpts := append([]jwt.ValidateOption{}, opts...)
		if rule != nil && len(rule.ValidateOptions) > 0 {
			combinedOpts = append(combinedOpts, rule.ValidateOptions...)
		}

		claims, err := verifyToken(o.verifier, payload, o.audience, o.replayChecker, combinedOpts)
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

func (o *OauthInterceptor) matchMethod(method string) *MethodRule {
	for _, rule := range o.methodRules {
		if matchPattern(method, rule.Pattern) {
			return &rule
		}
	}
	return nil
}

func (o *OauthInterceptor) extractTokenAndScheme(md metadata.MD) (scheme string, payload string, err error) {
	values := md.Get(o.tokenHeader)
	if len(values) == 0 {
		return "", "", errors.New("missing token header: " + o.tokenHeader)
	}

	raw := strings.TrimSpace(values[0])
	parts := strings.SplitN(raw, " ", 2)
	if len(parts) == 1 {
		return strings.ToLower(o.tokenScheme), parts[0], nil
	}

	return strings.ToLower(parts[0]), strings.TrimSpace(parts[1]), nil
}

// === Pattern matcher with "*" suffix support ===

func matchPattern(s, pattern string) bool {
	if prefix, ok := strings.CutSuffix(pattern, "*"); ok {
		return strings.HasPrefix(s, prefix)
	}
	return s == pattern
}
