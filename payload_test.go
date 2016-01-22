package slackcmd

import (
	"net/url"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPayload(t *testing.T) {
	assert := assert.New(t)

	params := url.Values{}
	params.Set("token", "1234567890")
	params.Set("command", "/ticket")
	params.Set("channel_name", "privategroup")

	payload := newPayloadByForm(params)
	assert.NotNil(payload)
	assert.Equal("1234567890", payload.Token)
	assert.Equal("/ticket", payload.Command)
	assert.False(payload.IsValid())
	assert.True(payload.IsPrivateGroup())

	params = url.Values{}
	params.Set("token", "1234567890")
	params.Set("command", "/ticket")
	params.Set("channel_name", "general")
	params.Set("channel_id", "T12345")

	payload = newPayloadByForm(params)
	assert.NotNil(payload)
	assert.True(payload.IsValid())
	assert.False(payload.IsPrivateGroup())
}

func TestTokenValidator(t *testing.T) {
	viper.Set("slackcmd.tokens", map[string]string{"ticket": "94972394739423492"})
	validator := NewTokenValidator()

	err := validator.Validate(&Payload{Token: "abcelksjdfldsf", Command: "/ticket"})
	assert.Equal(t, ErrTokenInvalid, err)

	err = validator.Validate(&Payload{Token: "94972394739423492", Command: "/ticket"})
	assert.Nil(t, err)
}
