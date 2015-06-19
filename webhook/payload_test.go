package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPayload(t *testing.T) {
	assert := assert.New(t)

	attach := NewAttachment("this is attachment text", "this is title", "http://example.com/title-link")
	payload := NewPayload()
	payload.AddAttachment(attach)
	assert.Len(payload.Attachments, 1)
}
