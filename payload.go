package slackcmd

import (
	"errors"
	"net/url"

	"github.com/spf13/viper"
)

// Payload contains all information that Slack post to server
type Payload struct {
	Token       string
	TeamID      string
	TeamDomain  string
	ChannelID   string
	ChannelName string
	UserID      string
	UserName    string
	Command     string
	Text        string
}

// IsPrivateGroup return true if command comes from a private group
func (p *Payload) IsPrivateGroup() bool {
	return p.ChannelName == "privategroup"
}

// IsValid return true if payload is valid, otherwise return false;
// Required fields: Token, Command, ChannelName, ChannelID.
func (p *Payload) IsValid() bool {
	return p.Token != "" &&
		p.ChannelName != "" &&
		p.ChannelID != "" &&
		p.Command != ""
}

// newPayloadByForm create payload from post/get form from request
func newPayloadByForm(form url.Values) *Payload {
	return &Payload{
		Token:       form.Get("token"),
		TeamID:      form.Get("team_id"),
		TeamDomain:  form.Get("team_domain"),
		ChannelID:   form.Get("channel_id"),
		ChannelName: form.Get("channel_name"),
		UserID:      form.Get("user_id"),
		UserName:    form.Get("user_name"),
		Command:     form.Get("command"),
		Text:        form.Get("text"),
	}
}

// PayloadValidator define interface which will be used to verify payload
type PayloadValidator interface {
	// Validate will check if given token (sent by Slack) is same as registered token for command
	Validate(payload *Payload) error
}

// NewTokenValidator return an implement of PayloadValidator,
// token validator will check if token sent by Slack Commands
// are match with token that in configuration
func NewTokenValidator() PayloadValidator {
	return &tokenValidator{}
}

var (
	// ErrTokenInvalid return when validate token failed
	ErrTokenInvalid = errors.New("token is invalid")
)

// tokenValidator used to validate token of Slack payload
type tokenValidator struct{}

// Validate token of Slack payload with configured token in app
func (t *tokenValidator) Validate(payload *Payload) error {
	tokens := viper.GetStringMapString("slackcmd.tokens")
	command := payload.Command[1:]

	if token, found := tokens[command]; found {
		if token != payload.Token {
			return ErrTokenInvalid
		}
	}

	return nil
}
