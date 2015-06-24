package webhook

import (
	"errors"
	"regexp"
	"strings"
)

const (
	// AttachmentColorGood define good color (green) for color fields of attachment
	AttachmentColorGood = "good"
	// AttachmentColorWarning define warning color (yellow) for color fields of attachment
	AttachmentColorWarning = "warning"
	// AttachmentColorDanger define danger color (red) for color fields of attachment
	AttachmentColorDanger = "danger"
)

var (
	// supportedColorName contains list of named color that are supported by Slack
	supportedColorName = map[string]bool{
		AttachmentColorGood:    true,
		AttachmentColorWarning: true,
		AttachmentColorDanger:  true,
	}

	// colorRegex to check if hex color is valid
	colorRegex = regexp.MustCompile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")

	// ErrColorInvalid return when set invalid color to attachment
	ErrColorInvalid = errors.New("inputed color is invalid, it should be 'good', 'warning', 'danger' or a hex color")
)

// Field represent a field object in attachment fields array
// see https://api.slack.com/docs/attachments
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// NewField create new Field with given params
func NewField(title, value string) *Field {
	return &Field{
		Title: title,
		Value: value,
	}
}

// Attachment represent an attachment object
// see https://api.slack.com/docs/attachments
type Attachment struct {
	Fallback   string   `json:"fallback,omitempty"`
	Color      string   `json:"color,omitemtpy"`
	PreText    string   `json:"pretext,omitemtpy"`
	AuthorName string   `json:"author_name,omitemtpy"`
	AuthorLink string   `json:"author_link,omitemtpy"`
	AuthorIcon string   `json:"author_icon,omitemtpy"`
	Title      string   `json:"title,omitemtpy"`
	TitleLink  string   `json:"title_link,omitemtpy"`
	Text       string   `json:"text,omitemtpy"`
	ImageURL   string   `json:"image_url,omitemtpy"`
	ThumbURL   string   `json:"thumb_url,omitemtpy"`
	Fields     []*Field `json:"fields,omitemtpy"`
	MrkDwnIn   []string `json:"mrkdwn_in"`
}

// NewAttachment create new attachment object, with good color by default
func NewAttachment(text, title, titleURL string) *Attachment {
	attachment := &Attachment{}
	attachment.Text = text
	attachment.Title = title
	attachment.TitleLink = titleURL

	attachment.SetColorToGood()
	attachment.setDefaultMrkDwnIn()
	return attachment
}

func (a *Attachment) setDefaultMrkDwnIn() {
	a.MrkDwnIn = []string{"pretext", "text", "fields"}
}

// AddField will append an attachment field to list of fields
func (a *Attachment) AddField(title, value string) {
	a.Fields = append(a.Fields, NewField(title, value))
}

// AddShortField will append an attachment field with short is true to list of fields
func (a *Attachment) AddShortField(title, value string) {
	field := NewField(title, value)
	field.Short = true
	a.Fields = append(a.Fields, field)
}

// SetColor is used to set color to field color of attachment
// color can 'good', 'warning', 'danger' or a hex color
func (a *Attachment) SetColor(color string) error {
	if strings.HasPrefix(color, "#") {
		if !colorRegex.MatchString(color) {
			return ErrColorInvalid
		}
	} else {
		if _, found := supportedColorName[color]; !found {
			return ErrColorInvalid
		}
	}

	a.Color = color
	return nil
}

// SetColorToGood will set color to green color
func (a *Attachment) SetColorToGood() {
	a.SetColor(AttachmentColorGood)
}

// SetColorToWarning will set color to yellow color
func (a *Attachment) SetColorToWarning() {
	a.SetColor(AttachmentColorWarning)
}

// SetColorToDanger will set color to red color
func (a *Attachment) SetColorToDanger() {
	a.SetColor(AttachmentColorDanger)
}
