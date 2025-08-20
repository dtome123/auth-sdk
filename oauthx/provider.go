// Module: oauthx — Unified OAuth/OIDC library for Go
// Providers: Google, Facebook, GitHub, Discord, Apple
// API shape: Authenticator{ AuthCodeURL, Exchange, FetchUser }
// Helpers: PKCE (S256), state, nonce, Apple client_secret (ES256)

package oauthx

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

type Provider string

const (
	Google   Provider = "google"
	Facebook Provider = "facebook"
	GitHub   Provider = "github"
	Discord  Provider = "discord"
	Apple    Provider = "apple"
)

// Config holds credentials for all providers. Empty credentials => provider disabled.
type Config struct {
	google   GoogleConfig
	facebook FacebookConfig
	github   GitHubConfig
	discord  DiscordConfig
	apple    AppleConfig
}

// User is a normalized user profile across providers.
type User struct {
	ID            string      `json:"id"`
	Email         string      `json:"email,omitempty"`
	EmailVerified *bool       `json:"email_verified,omitempty"`
	Name          string      `json:"name,omitempty"`
	AvatarURL     string      `json:"avatar_url,omitempty"`
	Raw           interface{} `json:"raw,omitempty"`
}

// TokenSafe returns a safe-to-log view of oauth2.Token (masked).
type TokenSafe struct {
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	IDToken      string    `json:"id_token,omitempty"`
}

// Authenticator is implemented per provider.
type Authenticator interface {
	Provider() Provider
	AuthCodeURLWithState(state, nonce, codeChallenge string) string
	ExchangeWithVerifier(ctx context.Context, code, codeVerifier string) (*oauth2.Token, error)
	AuthCodeURLNoState() string
	ExchangeNoVerifier(ctx context.Context, code string) (*oauth2.Token, error)
	FetchUser(ctx context.Context, tok *oauth2.Token) (User, error)
}

func MaskToken(t *oauth2.Token) TokenSafe {
	var idTok string
	if x, ok := t.Extra("id_token").(string); ok {
		idTok = truncate(x)
	}
	return TokenSafe{
		AccessToken:  truncate(t.AccessToken),
		RefreshToken: truncate(t.RefreshToken),
		TokenType:    t.TokenType,
		Expiry:       t.Expiry,
		IDToken:      idTok,
	}
}
