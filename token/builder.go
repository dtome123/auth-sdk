package token

import (
	"time"

	"github.com/dtome123/auth-sdk/jwt"
)

// ======================
// JWT Token Builder
// ======================

type builder struct {
	claims     jwt.Claims
	accessTTL  time.Duration
	refreshTTL time.Duration
	signer     jwt.Signer
}

// NewTokenBuilder creates a new builder for BearerToken.
func NewTokenBuilder(claims jwt.Claims, accessTTL, refreshTTL time.Duration, signer jwt.Signer) jwt.TokenBuilder {
	return &builder{
		claims:     claims,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		signer:     signer,
	}
}

// Build generates both access and refresh tokens.
func (b *builder) Build() (jwt.BearerToken, error) {
	accessToken, err := b.signer.Sign(b.claims, b.accessTTL)
	if err != nil {
		return jwt.BearerToken{}, err
	}

	refreshToken, err := b.signer.Sign(b.claims, b.refreshTTL)
	if err != nil {
		return jwt.BearerToken{}, err
	}

	return jwt.BearerToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int32(b.accessTTL.Seconds()),
		ExpiresAt:    time.Now().Add(b.accessTTL),
	}, nil
}
