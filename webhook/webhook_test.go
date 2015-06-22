package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)

	var err error
	err = SendMessage(nil)
	assert.Error(err)

	attach := NewAttachment("this is attachment text", "this is title", "http://example.com/title-link")
	payload := NewPayload()
	payload.AddAttachment(attach)

	err = SendMessage(payload)
	assert.Error(err)

	err = SendMessageToHook(payload, "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	assert.NoError(err)
}
