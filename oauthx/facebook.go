package oauthx

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
)

type facebookAuth struct {
	auth
}

type FacebookConfig struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
}

func (c *FacebookConfig) Validate() error {
	if c.ClientID == "" || c.ClientSecret == "" || c.RedirectURL == "" {
		return fmt.Errorf("provider %s config invalid", Facebook)
	}
	return nil
}

func NewFacebook(cfg FacebookConfig) (Authenticator, error) {
	return &facebookAuth{auth: auth{
		oc: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"public_profile", "email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.facebook.com/v19.0/dialog/oauth",
				TokenURL: "https://graph.facebook.com/v19.0/oauth/access_token",
			},
		},
	}}, nil
}

func (f *facebookAuth) Provider() Provider { return Facebook }
func (f *facebookAuth) AuthCodeURLWithState(state, nonce string, challenge string) string {
	return f.authCodeURL(state, nonce, challenge)
}
func (f *facebookAuth) ExchangeWithVerifier(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	return f.exchange(ctx, code, verifier)
}

func (f *facebookAuth) FetchUser(ctx context.Context, tok *oauth2.Token) (User, error) {
	cli := oauth2.NewClient(ctx, oauth2.StaticTokenSource(tok))
	resp, err := cli.Get("https://graph.facebook.com/v19.0/me?fields=id,name,email,picture.width(256).height(256)")
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()
	var fb struct {
		ID, Name, Email string
		Picture         struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&fb); err != nil {
		return User{}, err
	}
	return User{ID: fb.ID, Email: fb.Email, Name: fb.Name, AvatarURL: fb.Picture.Data.URL, Raw: fb}, nil
}
