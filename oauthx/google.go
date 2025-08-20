package oauthx

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleAuth struct {
	auth
}

type GoogleConfig struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
}

func (c *GoogleConfig) Validate() error {
	if c.ClientID == "" || c.ClientSecret == "" || c.RedirectURL == "" {
		return fmt.Errorf("provider %s config invalid", Google)
	}
	return nil
}

func NewGoogle(cfg GoogleConfig) (Authenticator, error) {
	oc := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
	prov, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, err
	}
	ver := prov.Verifier(&oidc.Config{ClientID: cfg.ClientID})
	return &googleAuth{auth: auth{oc: oc, verifier: ver}}, nil
}

func (g *googleAuth) Provider() Provider { return Google }
func (g *googleAuth) AuthCodeURLWithState(state, nonce, challenge string) string {
	return g.authCodeURL(state, nonce, challenge)
}
func (g *googleAuth) ExchangeWithVerifier(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	return g.exchange(ctx, code, verifier)
}

func (g *googleAuth) FetchUser(ctx context.Context, tok *oauth2.Token) (User, error) {
	rawID, _ := tok.Extra("id_token").(string)
	idTok, err := g.verifier.Verify(ctx, rawID)
	if err != nil {
		return User{}, err
	}
	var c struct {
		Sub, Email, Name, Picture string
		EmailVerified             *bool
	}
	if err := idTok.Claims(&c); err != nil {
		return User{}, err
	}
	return User{ID: c.Sub, Email: c.Email, EmailVerified: c.EmailVerified, Name: c.Name, AvatarURL: c.Picture, Raw: c}, nil
}
