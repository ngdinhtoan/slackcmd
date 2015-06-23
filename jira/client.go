package jira

import (
	jc "github.com/ngdinhtoan/go-jira-client"
	"github.com/spf13/viper"
)

var jiraClient *jc.Jira

func getJiraClient() *jc.Jira {
	if jiraClient == nil {
		jiraClient = jc.NewJira(
			viper.GetString("jira.host"),
			viper.GetString("jira.api_path"),
			viper.GetString("jira.activity_path"),
			&jc.Auth{
				Login:    viper.GetStringMapString("jira.auth")["login"],
				Password: viper.GetStringMapString("jira.auth")["password"],
			},
		)
	}

	return jiraClient
}

// getIssueDetail return issue by given issue ID
func getIssueDetail(issueID string) jc.Issue {
	return getJiraClient().Issue(issueID)
}

func getMarkdownUsername(user *jc.User) string {
	if user == nil {
		return ""
	}

	var name string
	if name = user.DisplayName; name == "" {
		name = user.Name
	}

	if user.Active {
		return name
	}

	return "_" + name + " (inactive)_"
}
