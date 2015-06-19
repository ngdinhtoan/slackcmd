package webhook

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

// SendMessageToHook will send a message to Slack via Incoming WebHooks to specific web hook
func SendMessageToHook(payload *Payload, webHookURL string) error {
	payloadData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Set("payload", string(payloadData))

	_, err = http.PostForm(webHookURL, params)
	return err
}

// SendMessage will send a message to Slack via Incoming WebHooks to configured web hook URL
func SendMessage(payload *Payload) error {
	webHookURL := viper.GetString("webhook.incomming_url")
	if webHookURL == "" {
		return errors.New("incomming webhook URL is not configured")
	}

	return SendMessageToHook(payload, webHookURL)
}
