package jwt

import "time"

// Signer defines the basic interface for signing JWT tokens.
type Signer interface {
	Sign(claims Claims, expiry time.Duration) (string, error)
}

// JWTSigner defines an extended interface for signing JWT tokens with algorithm metadata.
type JWTSigner interface {
	Sign(claims Claims, expiry time.Duration) (string, error)
	Algorithm() string
}

// Verifier defines the basic interface for verifying JWT tokens.
type Verifier interface {
	Verify(tokenStr string) (Claims, error)
}

// JWTVerifier defines an interface for verifying JWT tokens.
type JWTVerifier interface {
	Verify(tokenStr string) (Claims, error)
}

// BearerToken contains both access and refresh tokens and expiry metadata.
type BearerToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int32
	ExpiresAt    time.Time
}

// TokenBuilder defines the interface for generating JWT bearer tokens.
type TokenBuilder interface {
	Build() (BearerToken, error)
}
