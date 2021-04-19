package twitchAPI

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	clientID = "bug2jjuz0dkbmt1ybs1af6kh1jm968"
	// Consider storing the secret in an environment variable or a dedicated storage system.
	clientSecret = "f3yeu9mojrxchu6c2pxtw0lce6d7dt"
	oauth2Config *clientcredentials.Config
)

func GetTwitchToken() string {
	oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access token: %s\n", token.AccessToken)
	return GetTwitchToken()
}
