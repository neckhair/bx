package bexio

import (
	"context"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

const (
	authURL  = "https://idp.bexio.com/authorize"
	tokenURL = "https://idp.bexio.com/token"
	scopes   = "offline_access,contact_show,project_show"
)

type Client struct {
	tokenSource oauth2.TokenSource
	httpClient  *http.Client
}

func NewConfig(clientID, clientSecret string) oauth2.Config {
	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		Scopes: strings.Split(scopes, ","),
	}
}

func NewClient(tokenSource oauth2.TokenSource) *Client {
	return &Client{tokenSource: tokenSource, httpClient: http.DefaultClient}
}

func (c *Client) Do(ctx context.Context, url string) (*http.Response, error) {
	httpClient := oauth2.NewClient(ctx, c.tokenSource)
	req, _ := http.NewRequest("GET", "https://api.bexio.com/2.0/contact", nil)
	req.Header.Add("Accept", "application/json")
	return httpClient.Do(req)
}
