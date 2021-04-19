package twitchAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ENDPOINT_URL = "https://api.twitch.tv/helix/streams/?user_login="

func GetTwitchChannelStatus(twitchLogin string) string {

	// var TOKEN = GetTwitchToken()
	req, _ := http.NewRequest(http.MethodGet, ENDPOINT_URL+twitchLogin, nil)
	req.Header.Add("Client-ID", "bug2jjuz0dkbmt1ybs1af6kh1jm968")
	req.Header.Add("Authorization", "Bearer q50rgndfp6pe1ogilj31q3h4eokj9v")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	return (string(responseData))
}
