package jira

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	jc "github.com/ngdinhtoan/go-jira-client"
	"github.com/ngdinhtoan/slackcmd"
	"github.com/ngdinhtoan/slackcmd/webhook"
	"github.com/spf13/viper"
)

func init() {
	slackcmd.Register(&ticket{})
}

type ticket struct{}

// test if jira implement all function of slackcmd.Commander interface
var _ slackcmd.Commander = (*ticket)(nil)

var (
	errFetchIssue  = errors.New("can not fetch issue data")
	errWrongSyntax = errors.New("No ticket given, use `/ticket ticket_id [ticket_id]...`")
)

// GetCommand return list of command name that can be served by this commander
func (t *ticket) GetCommand() []string {
	return []string{"/ticket"}
}

// Execute when Slack sends command
func (t *ticket) Execute(payload *slackcmd.Payload, w http.ResponseWriter) error {
	issueIDs := strings.Fields(payload.Text)
	if len(issueIDs) == 0 {
		return errWrongSyntax
	}

	issueMap := map[string]bool{}
	for i := range issueIDs {
		issueMap[issueIDs[i]] = true
	}

	whPayload := webhook.NewPayload()
	whPayload.Text = fmt.Sprintf("`%s %s`", payload.Command, payload.Text)
	whPayload.Channel = payload.ChannelID

	wg := &sync.WaitGroup{}
	lk := &sync.Mutex{}
	notFoundIssues := make([]string, 0, len(issueMap))

	for issueID := range issueMap {
		wg.Add(1)

		go func(issueID string) {
			defer wg.Done()

			if attach, _ := t.fetchIssue(issueID); attach != nil {
				lk.Lock()
				defer lk.Unlock()

				issueMap[issueID] = true
				whPayload.AddAttachment(attach)
			} else {
				lk.Lock()
				defer lk.Unlock()

				notFoundIssues = append(notFoundIssues, issueID)
			}
		}(issueID)
	}

	wg.Wait()

	go webhook.SendMessage(whPayload)

	var err error
	if len(notFoundIssues) > 0 {
		err = fmt.Errorf("could not load information for issue: %s", strings.Join(notFoundIssues, ", "))
	}

	return err
}

func (t *ticket) fetchIssue(issueID string) (attach *webhook.Attachment, err error) {
	issue := getIssueDetail(issueID)
	if issue.Id == "" {
		err = errFetchIssue
		return
	}

	fn := getPrepareAttachmentFunc()
	attach = fn(&issue)

	return
}

// PrepareAttachmentFunc will create attachment for incomming webhook payload
type PrepareAttachmentFunc func(issue *jc.Issue) *webhook.Attachment

var defaultPrepareAttachmentFunc PrepareAttachmentFunc = func(issue *jc.Issue) *webhook.Attachment {
	issueURL := viper.GetString("jira.host") + "/browse/" + issue.Key

	attach := webhook.NewAttachment(issue.Fields.Description, issue.Fields.Summary, issueURL)
	attach.PreText = fmt.Sprintf("Ticket <%s|%s> is *%s*", issueURL, issue.Key, issue.Fields.Status.Name)
	if issue.Fields.Resolution != nil {
		attach.PreText += fmt.Sprintf(" with *%s* resolution", issue.Fields.Resolution.Name)
	}
	attach.AddShortField("Priority", issue.Fields.Priority.Name)
	attach.AddShortField("Status", issue.Fields.Status.Name)
	attach.AddShortField("Assignee", getMarkdownUsername(issue.Fields.Assignee))
	attach.AddShortField("Reporter", getMarkdownUsername(issue.Fields.Reporter))

	colorByPriority := viper.GetStringMapString("jira.priority_color")
	if color, found := colorByPriority[issue.Fields.Priority.Name]; found {
		attach.SetColor(color)
	}

	return attach
}

var customPrepareAttachmentFunc PrepareAttachmentFunc

func getPrepareAttachmentFunc() PrepareAttachmentFunc {
	if customPrepareAttachmentFunc != nil {
		return customPrepareAttachmentFunc
	}

	return defaultPrepareAttachmentFunc
}

// SetPrepareAttachmentFunc allow to set a custom PrepareAttachmentFunc
func SetPrepareAttachmentFunc(fn PrepareAttachmentFunc) {
	customPrepareAttachmentFunc = fn
}
