package oauthx

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
)

type discordAuth struct {
	auth
}

type DiscordConfig struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
}

func (c *DiscordConfig) Validate() error {
	if c.ClientID == "" || c.ClientSecret == "" || c.RedirectURL == "" {
		return fmt.Errorf("provider %s config invalid", Discord)
	}
	return nil
}

func NewDiscord(cfg DiscordConfig) (Authenticator, error) {
	return &discordAuth{auth: auth{
		oc: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"identify", "email"},
			Endpoint:     oauth2.Endpoint{AuthURL: "https://discord.com/api/oauth2/authorize", TokenURL: "https://discord.com/api/oauth2/token"},
		},
	}}, nil
}

func (d *discordAuth) Provider() Provider { return Discord }
func (d *discordAuth) AuthCodeURLWithState(state, nonce string, challenge string) string {
	return d.authCodeURL(state, nonce, challenge)
}
func (d *discordAuth) ExchangeWithVerifier(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	return d.exchange(ctx, code, verifier)
}
func (d *discordAuth) FetchUser(ctx context.Context, tok *oauth2.Token) (User, error) {
	cli := oauth2.NewClient(ctx, oauth2.StaticTokenSource(tok))
	resp, err := cli.Get("https://discord.com/api/users/@me")
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()
	var du struct{ ID, Username, GlobalName, Discriminator, Avatar, Email string }
	if err := json.NewDecoder(resp.Body).Decode(&du); err != nil {
		return User{}, err
	}
	avatarURL := ""
	if du.Avatar != "" {
		avatarURL = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", du.ID, du.Avatar)
	}
	name := du.GlobalName
	if name == "" {
		name = du.Username
	}
	return User{ID: du.ID, Email: du.Email, Name: name, AvatarURL: avatarURL, Raw: du}, nil
}
