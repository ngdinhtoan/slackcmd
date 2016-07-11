package webhook

// Payload define payload structure for Slack incoming webhook
type Payload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitemtpy"`
	Attachments []*Attachment `json:"attachments,omitemtpy"`
}

// NewPayload create new payload with default user name and icon emoji
func NewPayload() *Payload {
	return &Payload{
		Username:  "SlackCmd",
		IconEmoji: ":slack:",
	}
}

// AddAttachment append given attachment to payload attachment list
func (p *Payload) AddAttachment(attach *Attachment) {
	if attach == nil {
		return
	}

	p.Attachments = append(p.Attachments, attach)
}
