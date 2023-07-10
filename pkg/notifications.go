package pkg

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"
)

type SlackNotifyInput struct {
	ProjectName string
	ScriptName  string
	HistoryLink string
	Success     bool
}

func notifySlack(ctx context.Context, input SlackNotifyInput, webhookURL string) error {
	a := slack.Attachment{}
	a.Footer = input.HistoryLink
	a.AuthorName = fmt.Sprintf("%v | %v", input.ProjectName, input.ScriptName)
	a.Title = "Script Run Complete"
	if input.Success {
		a.Color = "good"
		a.Fallback = fmt.Sprintf("%v | %v : Run Successful", input.ProjectName, input.ScriptName)
		a.Text = fmt.Sprintf("Script %v Run Successful", input.ScriptName)
	} else {
		a.Color = "danger"
		a.Fallback = fmt.Sprintf("%v | %v : Run Failed", input.ProjectName, input.ScriptName)
		a.Text = fmt.Sprintf("Script %v Run Failed", input.ScriptName)
	}

	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{a},
	}
	return slack.PostWebhookContext(ctx, webhookURL, &msg)
}
