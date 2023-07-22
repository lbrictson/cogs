package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/cogs/ent"
	"github.com/slack-go/slack"
	"net/http"
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

type WebhookNotificationPayload struct {
	ProjectName     string            `json:"project_name"`
	ScriptName      string            `json:"script_name"`
	HistoryLink     string            `json:"history_link"`
	Success         bool              `json:"success"`
	TriggeredBy     string            `json:"triggered_by"`
	Trigger         string            `json:"trigger"`
	DurationSeconds int               `json:"duration_seconds"`
	Arguments       map[string]string `json:"arguments,omitempty"`
	RunID           int               `json:"run_id"`
}

func notifyWebhook(ctx context.Context, runID int, webhookURL string, db *ent.Client) error {
	history, err := getHistoryByID(ctx, db, runID)
	if err != nil {
		return err
	}
	script, err := getScriptByID(ctx, db, history.ScriptID)
	if err != nil {
		return err
	}
	project, err := getProjectByID(ctx, db, script.ProjectID)
	if err != nil {
		return err
	}
	input := WebhookNotificationPayload{
		ProjectName: project.Name,
		ScriptName:  script.Name,
		HistoryLink: fmt.Sprintf("%v/projects/%v/%v/history/%v",
			globalCallbackURL, script.ProjectID, script.ID, history.ID),
		Success:         history.Success,
		TriggeredBy:     history.TriggeredBy,
		Trigger:         history.Trigger,
		DurationSeconds: history.Duration,
		Arguments:       history.Arguments,
		RunID:           history.ID,
	}
	return doWebhookNotification(ctx, input, webhookURL)
}

func doWebhookNotification(ctx context.Context, input WebhookNotificationPayload, webhookURL string) error {
	d, err := json.Marshal(input)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
