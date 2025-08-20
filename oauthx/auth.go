package oauthx

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type auth struct {
	oc       *oauth2.Config
	verifier *oidc.IDTokenVerifier
}

func (a *auth) AuthCodeURLNoState() string {
	return a.oc.AuthCodeURL("")
}

func (a *auth) ExchangeNoVerifier(ctx context.Context, code string) (*oauth2.Token, error) {
	return a.oc.Exchange(ctx, code)
}

func (a *auth) authCodeURL(state, nonce, challenge string) string {
	opts := []oauth2.AuthCodeOption{}
	if challenge != "" {
		opts = append(opts,
			oauth2.SetAuthURLParam("code_challenge", challenge),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		)
	}
	if nonce != "" {
		opts = append(opts, oauth2.SetAuthURLParam("nonce", nonce))
	}
	return a.oc.AuthCodeURL(state, opts...)
}

func (a *auth) exchange(ctx context.Context, code, verifier string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	if opts == nil {
		opts = []oauth2.AuthCodeOption{}
	}
	if verifier != "" {
		opts = append(opts, oauth2.SetAuthURLParam("code_verifier", verifier))
	}
	return a.oc.Exchange(ctx, code, opts...)
}
