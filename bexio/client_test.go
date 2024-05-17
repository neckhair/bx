package bexio_test

import (
	"context"

	"github.com/neckhair/bx/bexio"
	"github.com/neckhair/bx/bexio/mocks"
)

func newTestClient(url string) *bexio.Client {
	tokenSource := &mocks.MockTokenSource{}
	client := bexio.NewClient(context.Background(), tokenSource)
	client.BaseUrl = url
	return client
}
