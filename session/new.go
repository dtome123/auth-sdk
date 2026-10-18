package session

import (
	"github.com/redis/go-redis/v9"
)

type Client[T any] struct {
	client redis.Cmdable
}

func New[T any](client redis.Cmdable) *Client[T] {
	return &Client[T]{
		client: client,
	}
}
