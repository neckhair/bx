package bexio

import (
	"context"
	"fmt"
	"log"

	"github.com/int128/oauth2cli"
	"github.com/int128/oauth2cli/oauth2params"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
)

var defaultBindAddresses []string = []string{"localhost:50424", "localhost:50710"}

func OAuthLogin(ctx context.Context, config oauth2.Config) (*oauth2.Token, error) {
	var token *oauth2.Token

	pkce, err := oauth2params.NewPKCE()
	if err != nil {
		return nil, err
	}

	ready := make(chan string, 1)
	defer close(ready)
	cfg := oauth2cli.Config{
		OAuth2Config:           config,
		AuthCodeOptions:        pkce.AuthCodeOptions(),
		TokenRequestOptions:    pkce.TokenRequestOptions(),
		LocalServerReadyChan:   ready,
		LocalServerBindAddress: defaultBindAddresses,
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case url := <-ready:
			log.Printf("Open %s", url)
			if err := browser.OpenURL(url); err != nil {
				log.Printf("could not open the browser: %s", err)
			}
			return nil
		case <-ctx.Done():
			return fmt.Errorf("context done while waiting for authorization: %w", ctx.Err())
		}
	})
	eg.Go(func() error {
		token, err = oauth2cli.GetToken(ctx, cfg)
		if err != nil {
			return fmt.Errorf("could not get a token: %w", err)
		}

		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, errors.Wrap(err, "authorization error")
	}

	return token, nil
}
