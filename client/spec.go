package client

import (
	"context"

	authPb "github.com/dtome123/auth-sdk/api/go/auth/v1"
	"github.com/dtome123/auth-sdk/types"
)

type Client interface {
	Sign(ctx context.Context, request *authPb.SignRequest) (*authPb.SignResponse, error)
	Refresh(ctx context.Context, request *authPb.RefreshRequest) (*authPb.RefreshResponse, error)
	Migration(ctx context.Context, seeding types.PermissionSeeding) error
}
