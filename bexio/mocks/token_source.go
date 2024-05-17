package mocks

import "golang.org/x/oauth2"

type MockTokenSource struct{}

func (ts *MockTokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{AccessToken: "fake-token"}
	return token, nil
}
