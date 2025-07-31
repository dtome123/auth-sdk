package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authPb "github.com/dtome123/auth-sdk/api/go/auth/v1"
	"github.com/dtome123/auth-sdk/jwtutils"
	"github.com/dtome123/auth-sdk/types"
)

const defaultAssertionTTL = 5 * time.Minute

type impl struct {
	endpoint     string
	iss          string
	sub          string
	authSvcAud   string
	assertionTTL time.Duration

	clientServiceSigner jwtutils.Signer
	clientOpts          []grpc.DialOption
	conn                *grpc.ClientConn
}

type ClientConfig struct {
	Endpoint string
	Iss      string
	Sub      string
	Aud      string
}

func NewClient(config ClientConfig, opts ...Option) (Client, error) {
	if config.Endpoint == "" {
		return nil, errors.New("endpoint is required")
	}

	c := &impl{
		endpoint:     config.Endpoint,
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

	if len(c.clientOpts) == 0 {
		c.clientOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	if c.clientServiceSigner == nil {
		return nil, errors.New("clientServiceSigner must not be nil")
	}

	// Open persistent gRPC connection
	conn, err := grpc.NewClient(config.Endpoint, c.clientOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}
	c.conn = conn

	return c, nil
}

// generateClientAssertion creates the signed JWT client assertion token.

func (c *impl) Sign(ctx context.Context, req *authPb.SignRequest) (*authPb.SignResponse, error) {
	clientAssertion, err := jwtutils.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	client := authPb.NewTokenServiceClient(c.conn)
	return client.Sign(ctx, req)
}

func (c *impl) Refresh(ctx context.Context, req *authPb.RefreshRequest) (*authPb.RefreshResponse, error) {
	clientAssertion, err := jwtutils.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return nil, err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	client := authPb.NewTokenServiceClient(c.conn)
	return client.Refresh(ctx, req)
}

func (c *impl) Migration(ctx context.Context, seeding types.PermissionSeeding) error {
	clientAssertion, err := jwtutils.GenerateClientAssertion(c.clientServiceSigner, c.iss, c.sub, c.authSvcAud, c.assertionTTL)
	if err != nil {
		return err
	}

	ctx = WithClientAssertion(ctx, clientAssertion)
	ctx = WithClientID(ctx, c.iss)

	client := authPb.NewAuthorizationServiceClient(c.conn)

	permissions := make([]*authPb.Permission, len(seeding.Permissions))
	for i, p := range seeding.Permissions {

		var impliedByActions []*authPb.ActionResource
		for _, action := range p.ImpliedByActions {
			impliedByActions = append(impliedByActions, &authPb.ActionResource{
				Resource: action.Resource,
				Action:   action.Action,
			})
		}

		permissions[i] = &authPb.Permission{
			Name:             p.Name,
			Description:      p.Description,
			Domain:           p.Domain,
			Resource:         p.Resource,
			Action:           p.Action,
			ImpliedByActions: impliedByActions,
		}
	}

	paths := make([]*authPb.PermissionPath, len(seeding.Paths))
	for i, p := range seeding.Paths {

		var scope authPb.RouteScope

		switch p.Type {
		case types.RouteScopePublic:
			scope = authPb.RouteScope_ROUTE_SCOPE_PUBLIC
		case types.RouteScopeAuthenticated:
			scope = authPb.RouteScope_ROUTE_SCOPE_AUTHENTICATED
		case types.RouteScopeAuthorized:
			scope = authPb.RouteScope_ROUTE_SCOPE_AUTHORIZED
		}

		paths[i] = &authPb.PermissionPath{
			Domain:   p.Domain,
			Path:     p.Path,
			Resource: p.Resource,
			Action:   p.Action,
			Type:     scope,
		}
	}

	_, err = client.MigratePermissions(ctx, &authPb.MigratePermissionsRequest{
		Permissions: permissions,
		Paths:       paths,
	})

	return err
}

func (c *impl) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
