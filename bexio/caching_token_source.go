package bexio

import (
	"context"

	"golang.org/x/oauth2"
)

type (
	ReadTokenFunc func() (*oauth2.Token, error)
	SaveTokenFunc func(*oauth2.Token) error
)

type CachingTokenSource struct {
	base          oauth2.TokenSource
	readTokenFunc ReadTokenFunc
	saveTokenFunc SaveTokenFunc
}

func NewCachingTokenSource(config *oauth2.Config, token *oauth2.Token, readFunc ReadTokenFunc, saveFunc SaveTokenFunc) (oauth2.TokenSource, error) {
	originalTokenSource := config.TokenSource(context.Background(), token)
	return oauth2.ReuseTokenSource(nil, &CachingTokenSource{
		base:          originalTokenSource,
		readTokenFunc: readFunc,
		saveTokenFunc: saveFunc,
	}), nil
}

// Tries to read from the given read function first. If that token is not present or not valid anymore,
// it reads from the given token source, which should renew the token automatically. After that, the
// token is saved via the given save token function.
func (c *CachingTokenSource) Token() (token *oauth2.Token, err error) {
	token, err = c.readTokenFunc()
	if err != nil {
		return nil, err
	}

	if token != nil && token.Valid() {
		return token, nil
	}

	if token, err = c.base.Token(); err != nil {
		return nil, err
	}

	err = c.saveTokenFunc(token)
	if err != nil {
		return token, err
	}

	return token, nil
}
