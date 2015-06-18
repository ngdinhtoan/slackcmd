package main

import (
	"log"

	"github.com/ngdinhtoan/slackcmd"
	_ "github.com/ngdinhtoan/slackcmd/jira"
	"github.com/spf13/viper"

	//	// feel free to custom issue attachment
	//	"github.com/ngdinhtoan/slackcmd/jira"
	//	"github.com/ngdinhtoan/slackcmd/webhook"
	//	jc "github.com/ngdinhtoan/go-jira-client"
)

func main() {
	log.Println("Loading configuration")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// viper.Debug()

	//	// feel free to custom issue attachment
	//	jira.SetPrepareAttachmentFunc(func(issue *jc.Issue) (*webhook.Attachment) {
	//		issueURL := viper.GetString("jira.host") + "/browse/" + issue.Key
	//
	//		attach := webhook.NewAttachment(issue.Fields.Description, issue.Fields.Summary, issueURL)
	//		attach.PreText = fmt.Sprintf("Ticket *<%s|%s>* is *%s*", issueURL, issue.Key, issue.Fields.Status.Name)
	//		attach.AddShortField("Priority", issue.Fields.Priority.Name)
	//		attach.AddShortField("Status", issue.Fields.Status.Name)
	//		attach.AddShortField("Assignee", issue.Fields.Assignee.DisplayName)
	//		attach.AddShortField("Reporter", issue.Fields.Reporter.DisplayName)
	//		attach.SetColor("#ffa")
	//
	//		return attach
	//	})

	log.Println("Starting server")
	slackcmd.StartServer("127.0.0.1", "12345", "/slackcmd")
}
