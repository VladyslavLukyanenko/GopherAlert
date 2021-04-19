package twitchAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ENDPOINT_URL = "https://api.twitch.tv/helix/streams/?user_login="
var TOKEN = GetTwitchToken()
var CLIENT_ID = clientID

func GetTwitchChannelStatus(twitchLogin string) string {

	req, _ := http.NewRequest(http.MethodGet, ENDPOINT_URL+twitchLogin, nil)
	req.Header.Add("Client-ID", CLIENT_ID)
	req.Header.Add("Authorization", "Bearer " + TOKEN)

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
