// Package webhook provide some useful structs and functions
// to send message to Slack via Incomming Webhook.
/*
Example send hello message to Slack via Incomming Webhook:

    package main

    import "github.com/ngdinhtoan/slackcmd/webhook"

    func main() {
        payload := webhook.NewPayload()
        payload.Text = "Hello Slack"
        payload.Username = "SlackCmd"

        webhook.SendMessageToHook(
            payload,
            "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
    }

Send rich message by Message Attachments (see https://api.slack.com/docs/attachments):

    package main

    import "github.com/ngdinhtoan/slackcmd/webhook"

    func main() {
        payload := webhook.NewPayload()
        payload.Text = "Hello Slack"
        payload.Username = "SlackCmd"

        attachment := webhook.NewAttachment("full message here", "title here", "http://title.url/here")
        attachment.SetColor("#F00")

        payload.AddAttachment(attachment)
        webhook.SendMessageToHook(
            payload,
            "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
    }

*/
package webhook
