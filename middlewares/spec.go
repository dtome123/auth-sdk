package middlewares

import (
	"context"

	"github.com/dtome123/auth-sdk/session"
)

// SessionGetter defines the generic contract for retrieving a session by userID and deviceID.
type SessionGetter[T any] interface {
	GetSession(ctx context.Context, userID, deviceID string) (*session.Session[T], error)
}
