package bexio

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
)

const (
	authURL    = "https://idp.bexio.com/authorize"
	tokenURL   = "https://idp.bexio.com/token"
	apiBaseURL = "https://api.bexio.com/2.0/"
	scopes     = "offline_access,contact_show,project_show"
)

type Client struct {
	tokenSource oauth2.TokenSource
	httpClient  *http.Client
}

type QueryParams map[string]string

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
	return &Client{tokenSource: oauth2.ReuseTokenSource(nil, tokenSource), httpClient: http.DefaultClient}
}

func (c *Client) Do(ctx context.Context, url string, query QueryParams) (*http.Response, error) {
	httpClient := oauth2.NewClient(ctx, c.tokenSource)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return httpClient.Do(req)
}

func (c *Client) Contacts(ctx context.Context, limit int) ([]Contact, error) {
	params := map[string]string{"limit": strconv.Itoa(limit)}
	resp, err := c.Do(ctx, apiBaseURL+"/contact", params)
	if err != nil {
		return nil, err
	}

	contacts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseListContactsResponse(contacts)
}
