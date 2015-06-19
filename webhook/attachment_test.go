package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttachment(t *testing.T) {
	assert := assert.New(t)

	attach := NewAttachment("this is attachment text", "this is title", "http://example.com/title-link")
	attach.AddField("description", "this is a description")
	attach.AddShortField("status", "open")

	assert.NotEmpty(attach.Text)
	assert.NotEmpty(attach.Title)
	assert.NotEmpty(attach.TitleLink)
	assert.Len(attach.Fields, 2)

	var err error
	err = attach.SetColor("#1234")
	assert.Equal(ErrColorInvalid, err)
	err = attach.SetColor("green")
	assert.Equal(ErrColorInvalid, err)
	err = attach.SetColor("#FAA")
	assert.NoError(err)
	err = attach.SetColor(AttachmentColorDanger)
	assert.NoError(err)
}
