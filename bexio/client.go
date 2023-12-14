package bexio

import (
	"context"
	"fmt"
	"net/http"
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
	BaseUrl     string
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
	return &Client{
		BaseUrl:     apiBaseURL,
		tokenSource: oauth2.ReuseTokenSource(nil, tokenSource),
		httpClient:  http.DefaultClient,
	}
}

func (c *Client) Get(ctx context.Context, url string, query QueryParams) (*http.Response, error) {
	httpClient := oauth2.NewClient(ctx, c.tokenSource)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode == http.StatusNotFound:
		return resp, NotFoundError
	case resp.StatusCode == http.StatusUnauthorized:
		return resp, UnauthorizedError
	case resp.StatusCode >= 300:
		return resp, fmt.Errorf("http error %s", resp.Status)
	default:
		return resp, nil
	}
}
