package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/neckhair/bx/bexio"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Reads the token from the given json file
func ReadTokenFromFile() (*oauth2.Token, error) {
	token := &oauth2.Token{}

	jsonFile, err := os.Open(TokenFileFullPath)
	if err != nil {
		return nil, errors.Wrap(err, "error loading token from file")
	}

	byteValue, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, token); err != nil {
		return nil, errors.Wrap(err, "auth token unmarshal failed")
	}

	return token, jsonFile.Close()
}

// Saves the token as json into the token file
func SaveTokenToFile(token *oauth2.Token) error {
	file, err := json.MarshalIndent(token, "", "")
	if err != nil {
		return errors.Wrap(err, "error saving token")
	}
	return os.WriteFile(TokenFileFullPath, file, 0o644)
}

func NewCachingTokenSource() (oauth2.TokenSource, error) {
	token, err := ReadTokenFromFile()
	if err != nil {
		return nil, err
	}

	cfg := bexio.NewConfig(Credentials())
	return bexio.NewCachingTokenSource(&cfg, token, ReadTokenFromFile, SaveTokenToFile)
}
