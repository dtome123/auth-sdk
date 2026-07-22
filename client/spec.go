package client

import (
	"context"

	authPb "github.com/dtome123/auth-sdk/api/auth/v1"
)

type Client interface {
	Sign(ctx context.Context, request *authPb.SignRequest) (*authPb.SignResponse, error)
	Refresh(ctx context.Context, request *authPb.RefreshRequest) (*authPb.RefreshResponse, error)
	Migration(ctx context.Context, request *authPb.MigratePermissionsRequest) error
}
