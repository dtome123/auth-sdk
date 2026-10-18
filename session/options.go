package session

// Option configures a generic session Client.
type Option[T any] func(*Client[T])

// WithKeyBuilder configures a custom key format builder function.
func WithKeyBuilder[T any](builder func(userID, deviceID string) string) Option[T] {
	return func(c *Client[T]) {
		if builder != nil {
			c.keyBuilder = builder
		}
	}
}
