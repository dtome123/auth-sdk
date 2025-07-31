package jwtutils

import (
	"context"
	"fmt"
	"time"

	authPb "github.com/dtome123/auth-sdk/api/go/auth/v1"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcTokenSource implements oauth2.TokenSource backed by gRPC to AuthService.
type GrpcTokenSource struct {
	authSvcEndpoint string
	signer          Signer
	iss             string
	sub             string
	aud             string
	assertionTTL    time.Duration
	opts            []grpc.DialOption
}

// NewGrpcTokenSource constructs a token source from gRPC AuthService.
func NewGrpcTokenSource(
	authSvcEndpoint string,
	iss, sub, aud string,
	signer Signer,
	assertTTL time.Duration,
	opts ...grpc.DialOption,
) *GrpcTokenSource {
	return &GrpcTokenSource{
		authSvcEndpoint: authSvcEndpoint,
		signer:          signer,
		iss:             iss,
		sub:             sub,
		aud:             aud,
		assertionTTL:    assertTTL,
		opts:            opts,
	}
}

// Token implements oauth2.TokenSource interface.
func (c *GrpcTokenSource) Token() (*oauth2.Token, error) {
	assertion, err := GenerateClientAssertion(c.signer, c.iss, c.sub, c.aud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	opts := c.opts
	if opts == nil {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	conn, err := grpc.NewClient(c.authSvcEndpoint, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}
	defer conn.Close()

	client := authPb.NewTokenServiceClient(conn)
	resp, err := client.Token(context.Background(), &authPb.TokenRequest{
		GrantType:           "client_credentials",
		ClientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
		ClientAssertion:     assertion,
	})
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}

	return &oauth2.Token{
		AccessToken: resp.AccessToken,
		TokenType:   resp.TokenType,
		Expiry:      time.Now().Add(time.Duration(resp.ExpiresIn) * time.Second),
	}, nil
}
