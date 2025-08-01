package middlewares

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/oauth2"
)

// oauth2PerRPCCredentials implements grpc.PerRPCCredentials using an OAuth2 token source.
type oauth2PerRPCCredentials struct {
	src                 oauth2.TokenSource
	requireTransportSec bool
	header              string // Header key (e.g. "authorization", lowercase)
	scheme              string // Token scheme (e.g. "bearer", lowercase or empty)
}

// Option defines functional options for customizing Oauth2PerRPCCredentials.
type Option func(*oauth2PerRPCCredentials)

// WithHeader sets a custom header name (normalized to lowercase).
func WithHeader(header string) Option {
	return func(o *oauth2PerRPCCredentials) {
		h := strings.ToLower(strings.TrimSpace(header))
		if h != "" {
			o.header = h
		}
	}
}

// WithScheme sets a custom scheme (normalized to lowercase).
func WithScheme(scheme string) Option {
	return func(o *oauth2PerRPCCredentials) {
		o.scheme = strings.ToLower(strings.TrimSpace(scheme))
	}
}

// NewOauth2PerRPCCredentials creates credentials using a token source.
// By default, uses "authorization" header with "bearer" scheme.
// Use options like WithHeader(...) and WithScheme(...) to customize.
func NewOauth2PerRPCCredentials(src oauth2.TokenSource, requireTLS bool, opts ...Option) *oauth2PerRPCCredentials {
	cred := &oauth2PerRPCCredentials{
		src:                 src,
		requireTransportSec: requireTLS,
		header:              "authorization",
		scheme:              "bearer",
	}
	for _, opt := range opts {
		opt(cred)
	}
	return cred
}

// RequireTransportSecurity indicates if TLS is required.
func (o *oauth2PerRPCCredentials) RequireTransportSecurity() bool {
	return o.requireTransportSec
}

// GetRequestMetadata injects the token into gRPC metadata using configured header and scheme.
func (o *oauth2PerRPCCredentials) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	token, err := o.src.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	if !token.Valid() {
		return nil, fmt.Errorf("token is invalid")
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("access token is empty")
	}

	header := strings.ToLower(strings.TrimSpace(o.header))
	if header == "" {
		header = "authorization" // fallback if somehow missing
	}

	scheme := strings.ToLower(strings.TrimSpace(o.scheme))

	value := token.AccessToken
	if scheme != "" {
		value = scheme + " " + value
	}

	return map[string]string{
		header: value,
	}, nil
}
