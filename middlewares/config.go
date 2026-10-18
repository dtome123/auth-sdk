package middlewares

import "github.com/dtome123/auth-sdk/session"

type SessionConfig[T any] struct {
	Repo          session.Repository[T]
	UserIDClaim   string
	DeviceIDClaim string
	JTIClaim      string
}
