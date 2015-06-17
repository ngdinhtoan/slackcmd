package slackcmd

import "net/url"

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
