package jwtutils

import (
	"time"
)

// BearerToken contains both access and refresh tokens and expiry metadata.
type BearerToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int32
	ExpiresAt    time.Time
}

// jwtTokenBuilder implements TokenBuilder using a JWT signer.
type jwtTokenBuilder struct {
	claims     Claims
	accessTTL  time.Duration
	refreshTTL time.Duration
	signer     Signer
}

// NewTokenBuilder creates a new builder for BearerToken.
func NewTokenBuilder(claims Claims, accessTTL, refreshTTL time.Duration, signer Signer) TokenBuilder {
	return &jwtTokenBuilder{
		claims:     claims,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		signer:     signer,
	}
}

// Build generates both access and refresh tokens.
func (b *jwtTokenBuilder) Build() (BearerToken, error) {
	accessToken, err := b.signer.Sign(b.claims, b.accessTTL)
	if err != nil {
		return BearerToken{}, err
	}

	refreshToken, err := b.signer.Sign(b.claims, b.refreshTTL)
	if err != nil {
		return BearerToken{}, err
	}

	return BearerToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int32(b.accessTTL.Seconds()),
		ExpiresAt:    time.Now().Add(b.accessTTL),
	}, nil
}
