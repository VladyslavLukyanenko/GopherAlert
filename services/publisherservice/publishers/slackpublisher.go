package publishers

import (
	"bytes"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PublishToSlack(webhook *core.Webhook) {
	log.Debugf("Received webhook into publisher: %s", webhook.URI)
	var jsonStr = []byte("{\"text\":\"Hello, World!\"}") //todo: Create an Embed and replace placeholders with data in webhook.JSON
	res, err := http.Post(webhook.URI, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Errorf("Error when sending request publish@slack: %s", err.Error())
		return
	}
	if res.StatusCode != 200 {
		log.Errorf("Status code not 200")
		return
	}
		log.Debugf("Sent webhook publish@slack")
}
