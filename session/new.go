package session

import (
	"github.com/redis/go-redis/v9"
)

// Client is a generic Redis-backed session repository for Session[T].
type Client[T any] struct {
	client     redis.Cmdable
	keyBuilder func(userID, deviceID string) string
}

// Ensure Client implements Repository interface.
var _ Repository[any] = (*Client[any])(nil)

// New creates a new generic session client.
func New[T any](client redis.Cmdable, opts ...Option[T]) *Client[T] {
	return NewClient[T](client, opts...)
}

// NewClient creates a new generic session client.
func NewClient[T any](client redis.Cmdable, opts ...Option[T]) *Client[T] {
	c := &Client[T]{
		client:     client,
		keyBuilder: BuildSessionKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// NewAuthenticationRepository creates a new generic session repository (backward compatibility alias).
func NewAuthenticationRepository[T any](client redis.Cmdable, opts ...Option[T]) *Client[T] {
	return NewClient[T](client, opts...)
}

// Redis returns the underlying Redis client.
func (repo *Client[T]) Redis() redis.Cmdable {
	return repo.client
}
