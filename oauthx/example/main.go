// ─────────────────────────────────────────────────────────────────────────────
// File: examples/server/main.go (usage example with Gin)
// This file shows how to use oauthx in a web server. Place it under examples/server in your repo.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dtome123/auth-sdk/oauthx"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}

	r, err := oauthx.NewRegistry(
		oauthx.WithGoogleProvider(oauthx.GoogleConfig{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  "http://localhost:8080/auth/google/callback",
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()
	app.GET("/", func(c *gin.Context) { c.String(200, "OK") })

	app.GET("/auth/google/login", func(c *gin.Context) {

		a, err := r.GetGoogleProvider()
		if err != nil {
			c.JSON(404, gin.H{"error": "provider not enabled"})
			return
		}
		state := oauthx.GenerateState()
		nonce := oauthx.GenerateNonce()
		verifier, challenge := oauthx.PKCEPair()
		c.SetCookie("state", state, 600, "/", "", false, true)
		c.SetCookie("nonce", nonce, 600, "/", "", false, true)
		c.SetCookie("verifier", verifier, 600, "/", "", false, true)
		c.Redirect(http.StatusFound, a.AuthCodeURLWithState(state, nonce, challenge))
	})

	app.GET("/auth/google/callback", func(c *gin.Context) {

		a, err := r.GetGoogleProvider()
		if err != nil {
			c.JSON(404, gin.H{"error": "provider not enabled"})
			return
		}
		expState, _ := c.Cookie("state")
		if c.Query("state") != expState {
			c.JSON(400, gin.H{"error": "state mismatch"})
			return
		}
		code := c.Query("code")
		if code == "" {
			c.JSON(400, gin.H{"error": "missing code"})
			return
		}
		verifier, _ := c.Cookie("verifier")
		tok, err := a.ExchangeWithVerifier(c, code, verifier)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		user, err := a.FetchUser(c, tok)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"provider": "google", "user": user, "token": oauthx.MaskToken(tok)})
	})

	log.Println("listening :8080")
	_ = app.Run(":8080")
}

func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
