package session

import "context"

// Repository defines the contract for generic session persistence.
type Repository[T any] interface {
	GetSession(ctx context.Context, userID, deviceID string) (*Session[T], error)
	UpsertSession(ctx context.Context, session Session[T]) error
	DeleteSession(ctx context.Context, userID, deviceID string) error
}
