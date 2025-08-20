package oauthx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

type appleAuth struct {
	auth
	cfg AppleConfig
}

type AppleConfig struct {
	RedirectURL string

	// Apple uses a JWT client_secret signed with your private key (ES256)
	ClientID  string // Service ID (web) or App ID (native)
	TeamID    string
	KeyID     string
	KeyPEM    string        // contents of AuthKey_XXXXXX.p8
	ClientTTL time.Duration // default 5m
}

func (c *AppleConfig) Validate() error {
	if c.ClientID == "" || c.TeamID == "" || c.KeyID == "" || c.KeyPEM == "" || c.RedirectURL == "" {
		return fmt.Errorf("provider %s config invalid", Apple)
	}

	if c.ClientTTL <= 0 {
		c.ClientTTL = 5 * time.Minute
	}

	if c.ClientTTL > 6*30*24*time.Hour {
		return fmt.Errorf("max client TTL is 6 months (180 days)")
	}

	return nil
}

// NewApple khởi tạo provider Apple OAuth2/OpenID
func NewApple(cfg AppleConfig) (Authenticator, error) {
	secret, err := makeAppleClientSecret(cfg)
	if err != nil {
		return nil, err
	}

	oc := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: secret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{"openid", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://appleid.apple.com/auth/authorize",
			TokenURL: "https://appleid.apple.com/auth/token",
		},
	}

	prov, err := oidc.NewProvider(context.Background(), "https://appleid.apple.com")
	if err != nil {
		return nil, err
	}
	ver := prov.Verifier(&oidc.Config{ClientID: cfg.ClientID})

	return &appleAuth{
		auth: auth{oc: oc, verifier: ver},
		cfg:  cfg,
	}, nil
}

func (a *appleAuth) Provider() Provider { return Apple }

func (a *appleAuth) AuthCodeURLWithState(state, nonce, challenge string) string {
	return a.authCodeURL(state, nonce, challenge)
}

func (a *appleAuth) ExchangeWithVerifier(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	secret, err := makeAppleClientSecret(a.cfg)
	if err != nil {
		return nil, err
	}
	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("client_id", a.cfg.ClientID),
		oauth2.SetAuthURLParam("client_secret", secret),
	}
	if verifier != "" {
		opts = append(opts, oauth2.SetAuthURLParam("code_verifier", verifier))
	}
	return a.oc.Exchange(ctx, code, opts...)
}

func (a *appleAuth) FetchUser(ctx context.Context, tok *oauth2.Token) (User, error) {
	rawID, _ := tok.Extra("id_token").(string)
	idTok, err := a.verifier.Verify(ctx, rawID)
	if err != nil {
		return User{}, err
	}

	var c struct {
		Sub           string
		Email         string
		EmailVerified *bool `json:"email_verified"`
	}
	if err := idTok.Claims(&c); err != nil {
		return User{}, err
	}

	return User{
		ID:            c.Sub,
		Email:         c.Email,
		EmailVerified: c.EmailVerified,
		Raw:           c,
	}, nil
}

func makeAppleClientSecret(cfg AppleConfig) (string, error) {
	if cfg.KeyPEM == "" || cfg.KeyID == "" || cfg.TeamID == "" || cfg.ClientID == "" {
		return "", errors.New("missing Apple credentials: APPLE_KEY_PEM, APPLE_KEY_ID, APPLE_TEAM_ID, APPLE_CLIENT_ID")
	}

	priv, err := jwt.ParseECPrivateKeyFromPEM([]byte(cfg.KeyPEM))
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"iss": cfg.TeamID,
		"iat": now.Unix(),
		"exp": now.Add(cfg.ClientTTL).Unix(),
		"aud": "https://appleid.apple.com",
		"sub": cfg.ClientID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t.Header["kid"] = cfg.KeyID
	return t.SignedString(priv)
}
