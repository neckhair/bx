package bexio

import (
	"context"
	"fmt"
	"io"
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
	BaseUrl    string
	httpClient *http.Client
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

func NewClient(ctx context.Context, tokenSource oauth2.TokenSource) *Client {
	return &Client{
		BaseUrl:    apiBaseURL,
		httpClient: oauth2.NewClient(ctx, oauth2.ReuseTokenSource(nil, tokenSource)),
	}
}

func (c *Client) Request(req *http.Request, query QueryParams) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
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

func (c *Client) Get(url string, query QueryParams) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	return c.Request(req, query)
}

func (c *Client) Post(url string, query QueryParams, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, body)
	return c.Request(req, query)
}
