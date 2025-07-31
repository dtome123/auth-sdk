package middlewares

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// Oauth2PerRPCCredentials implements grpc.PerRPCCredentials using an OAuth2 token source.
type Oauth2PerRPCCredentials struct {
	src                 oauth2.TokenSource // Source to obtain OAuth2 access tokens
	requireTransportSec bool               // Whether to enforce secure transport (TLS)
}

// NewOauth2PerRPCCredentials creates a new instance of Oauth2PerRPCCredentials.
// - src: the OAuth2 token source
// - requireTLS: if true, requires TLS for gRPC connections
func NewOauth2PerRPCCredentials(src oauth2.TokenSource, requireTLS bool) *Oauth2PerRPCCredentials {
	return &Oauth2PerRPCCredentials{
		src:                 src,
		requireTransportSec: requireTLS,
	}
}

// RequireTransportSecurity returns whether this credential requires a secure transport (TLS).
func (o *Oauth2PerRPCCredentials) RequireTransportSecurity() bool {
	return o.requireTransportSec
}

// GetRequestMetadata adds the "authorization" header with a bearer token to each gRPC request.
func (o *Oauth2PerRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
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

	return map[string]string{
		"authorization": "Bearer " + token.AccessToken,
	}, nil
}
