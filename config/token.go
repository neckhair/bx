package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const tokenKey = "token"

type OAuth2Token struct {
	Accesstoken  string    `mapstructure:"accesstoken"`
	Tokentype    string    `mapstructure:"tokentype"`
	Refreshtoken string    `mapstructure:"refreshtoken"`
	Expiry       time.Time `mapstructure:"expiry"`
}

type provider struct{}

var TokenProvider = &provider{}

func (p *provider) Token() (*oauth2.Token, error) {
	token, err := Token()
	if err != nil {
		return token, nil
	}
	return token, nil
}

func Token() (*oauth2.Token, error) {
	var token oauth2.Token

	err := viper.UnmarshalKey(tokenKey, &token)
	if err != nil {
		return nil, fmt.Errorf("cannot read token from config: %w", err)
	}

	return &token, nil
}

func SetToken(token *oauth2.Token) {
	viper.Set(tokenKey, token)
}
