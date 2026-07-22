package client

import (
	"context"
	"errors"
	"time"

	authPb "github.com/dtome123/auth-sdk/api/auth/v1"
	"github.com/dtome123/auth-sdk/jwt"
	"google.golang.org/grpc"
)

const defaultAssertionTTL = 5 * time.Minute

type impl struct {
	tokenClient  authPb.TokenServiceClient
	authClient   authPb.AuthorizationServiceClient
	iss          string
	sub          string
	authSvcAud   string
	assertionTTL time.Duration

	clientServiceSigner jwt.Signer
}

type ClientConfig struct {
	Iss string
	Sub string
	Aud string
}

func NewClient(config ClientConfig, conn *grpc.ClientConn, opts ...Option) (Client, error) {

	c := &impl{
		tokenClient:  authPb.NewTokenServiceClient(conn),
		authClient:   authPb.NewAuthorizationServiceClient(conn),
		iss:          config.Iss,
		sub:          config.Sub,
		authSvcAud:   config.Aud,
		assertionTTL: defaultAssertionTTL,
	}

	// Apply functional options
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.clientServiceSigner == nil {
		return nil, errors.New("clientServiceSigner must not be nil")
	}

	return c, nil
}

// generateClientAssertion creates the signed JWT client assertion token.

func (c *impl) Sign(ctx context.Context, req *authPb.SignRequest) (*authPb.SignResponse, error) {
	clientAssertion, err := jwt.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	return c.tokenClient.Sign(ctx, req)
}

func (c *impl) Refresh(ctx context.Context, req *authPb.RefreshRequest) (*authPb.RefreshResponse, error) {
	clientAssertion, err := jwt.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	return c.tokenClient.Refresh(ctx, req)
}

func (c *impl) Migration(ctx context.Context, req *authPb.MigratePermissionsRequest) error {
	clientAssertion, err := jwt.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	_, err = c.authClient.MigratePermissions(ctx, req)

	return err
}

func (c *impl) Close() error {
	return nil
}
