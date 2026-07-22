package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	authPb "github.com/dtome123/auth-sdk/api/auth/v1"
	"github.com/dtome123/auth-sdk/jwt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcTokenSourceConfig holds configuration parameters for GrpcTokenSource.
type GrpcTokenSourceConfig struct {
	AuthSvcEndpoint string
	Issuer          string
	Subject         string
	Audience        string
	Signer          jwt.Signer
	AssertionTTL    time.Duration
	DialOptions     []grpc.DialOption
}

// Validate checks whether the config struct has valid required fields.
func (cfg *GrpcTokenSourceConfig) Validate() error {
	if cfg.AuthSvcEndpoint == "" {
		return errors.New("authSvcEndpoint is required")
	}
	if cfg.Issuer == "" {
		return errors.New("issuer (iss) is required")
	}
	if cfg.Subject == "" {
		return errors.New("subject (sub) is required")
	}
	if cfg.Audience == "" {
		return errors.New("audience (aud) is required")
	}
	if cfg.Signer == nil {
		return errors.New("signer is required")
	}
	if cfg.AssertionTTL <= 0 {
		return errors.New("assertionTTL must be greater than zero")
	}
	return nil
}

// GrpcTokenSource implements oauth2.TokenSource backed by gRPC to AuthService.
type GrpcTokenSource struct {
	authSvcEndpoint string
	signer          jwt.Signer
	iss             string
	sub             string
	aud             string
	assertionTTL    time.Duration

	conn   *grpc.ClientConn
	client authPb.TokenServiceClient
}

// NewGrpcTokenSource constructs a token source from gRPC AuthService using GrpcTokenSourceConfig.
func NewGrpcTokenSource(cfg GrpcTokenSourceConfig) (*GrpcTokenSource, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid grpc token source config: %w", err)
	}

	dialOpts := cfg.DialOptions
	if len(dialOpts) == 0 {
		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	conn, err := grpc.NewClient(cfg.AuthSvcEndpoint, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &GrpcTokenSource{
		authSvcEndpoint: cfg.AuthSvcEndpoint,
		signer:          cfg.Signer,
		iss:             cfg.Issuer,
		sub:             cfg.Subject,
		aud:             cfg.Audience,
		assertionTTL:    cfg.AssertionTTL,
		conn:            conn,
		client:          authPb.NewTokenServiceClient(conn),
	}, nil
}

// Close closes the underlying gRPC connection if open.
func (c *GrpcTokenSource) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		c.client = nil
		return err
	}
	return nil
}

// Token implements oauth2.TokenSource interface.
func (c *GrpcTokenSource) Token() (*oauth2.Token, error) {
	assertion, err := jwt.GenerateClientAssertion(c.signer, c.iss, c.sub, c.aud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Token(context.Background(), &authPb.TokenRequest{
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
