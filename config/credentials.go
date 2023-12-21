package config

import (
	"github.com/spf13/viper"
)

const (
	clientIDKey     = "client_id"
	clientSecretKey = "client_secret"
)

func SetCredentials(clientID, clientSecret string) {
	viper.Set(clientIDKey, clientID)
	viper.Set(clientSecretKey, clientSecret)
}

func Credentials() (string, string) {
	return viper.GetString("client_id"), viper.GetString("client_secret")
}

func CredentialsPresent() bool {
	clientID, clientSecret := Credentials()
	return clientID != "" && clientSecret != ""
}
