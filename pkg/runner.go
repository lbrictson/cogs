package pkg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/lbrictson/cogs/ent"
	"os"
	"os/exec"
	"strings"
	"time"
)

type RunScriptInput struct {
	Script         ScriptModel
	Caller         string
	Trigger        string
	Args           map[string]string
	ProjectID      int
	ProjectName    string
	SuccessChannel *NotificationChannelModel
	FailureChannel *NotificationChannelModel
}

func runScript(ctx context.Context, db *ent.Client, input RunScriptInput) int {
	uu := uuid.New().String()
	i := CreateHistoryInput{
		ScriptID:        input.Script.ID,
		Success:         false,
		Output:          "",
		DurationSeconds: 0,
		Trigger:         input.Trigger,
		TriggeredBy:     input.Caller,
		RunID:           uu,
		ExitCode:        0,
		Arguments:       input.Args,
		Status:          "running",
	}
	history, err := createHistory(ctx, db, i)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return 0
	}
	go doScriptRun(ctx, db, input, uu, history.ID)
	return history.ID
}

func doScriptRun(ctx context.Context, db *ent.Client, input RunScriptInput, runID string, historyID int) {
	// Prior to running the script we fetch all the secrets for the project to inject as env vars
	secrets, err := getProjectSecrets(ctx, db, input.ProjectID)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	valuesToScrub := []string{}
	for _, secret := range secrets {
		valuesToScrub = append(valuesToScrub, secret.Value)
	}
	// Create a file with the script inside
	start := time.Now()
	input.Script.Script = strings.Replace(input.Script.Script, "\r\n", "\n", -1)
	os.MkdirAll(dataDirectory+"/scripts/"+runID, 0755)
	err = os.WriteFile(dataDirectory+"/scripts/"+runID+"/script.sh", []byte(input.Script.Script), 0755)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	exec.Command("chmod", "+x", dataDirectory+"/scripts/"+runID+"/script.sh").Run()
	var cancel context.CancelFunc
	scriptContext, cancel := context.WithTimeout(context.Background(), time.Duration(input.Script.TimeoutSeconds)*time.Second)
	defer cancel()
	// Run the script
	cmd := exec.CommandContext(scriptContext, "bash", "-c", dataDirectory+"/scripts/"+runID+"/script.sh >> tmp/"+runID+".log 2>&1")
	cmd.Env = os.Environ()
	for k, v := range input.Args {
		cmd.Env = append(cmd.Env, "COGS_"+strings.ToUpper(k)+"="+v)
	}
	for _, secret := range secrets {
		cmd.Env = append(cmd.Env, "COGS_SECRET_"+strings.ToUpper(secret.Name)+"="+secret.Value)
	}
	out, commandRunErr := cmd.CombinedOutput()
	if commandRunErr != nil {
		LogFromCtx(ctx).Error(commandRunErr.Error())
	}
	// Update the history
	update := UpdateHistoryInput{
		Success:         nil,
		Output:          nil,
		DurationSeconds: nil,
		Trigger:         nil,
		TriggeredBy:     nil,
		RunID:           nil,
		ExitCode:        nil,
		Status:          nil,
		Arguments:       nil,
	}
	t := true
	f := false
	runOutcomeSuccess := true
	if commandRunErr != nil {
		update.Success = &f
		runOutcomeSuccess = false
	} else {
		update.Success = &t
	}
	o := string(out)
	if runOutcomeSuccess != true {
		o = o + "\n" + commandRunErr.Error()
	}
	for _, x := range valuesToScrub {
		o = strings.ReplaceAll(o, x, "********")
	}
	// Read the file to get the full output
	fileContents, readFileErr := os.ReadFile("tmp/" + runID + ".log")
	if readFileErr == nil {
		o = o + "\n" + string(fileContents)
		for _, x := range valuesToScrub {
			o = strings.ReplaceAll(o, x, "********")
		}
		// Remove the log file
		os.Remove("tmp/" + runID + ".log")
	}
	update.Output = &o
	d := time.Since(start).Seconds()
	// convert to int
	di := int(d)
	update.DurationSeconds = &di
	status := "finished"
	update.Status = &status
	// Delete the file
	exec.Command("rm", "-rf", dataDirectory+"/scripts/"+runID).Run()
	// Return
	updateStatsInput := UpdateScriptStatsInput{
		ScriptID:         input.Script.ID,
		IncrementSuccess: false,
		IncrementError:   false,
		DurationSeconds:  di,
	}
	_, err = updateHistory(ctx, db, historyID, update)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	if runOutcomeSuccess {
		updateStatsInput.IncrementSuccess = true
	} else {
		updateStatsInput.IncrementError = true
	}
	updateScriptStats(ctx, db, updateStatsInput)
	now := time.Now()
	if runOutcomeSuccess {
		if input.SuccessChannel != nil {
			switch input.SuccessChannel.Type {
			case "slack":
				notificationErr := notifySlack(ctx, SlackNotifyInput{
					ProjectName: input.ProjectName,
					ScriptName:  input.Script.Name,
					HistoryLink: fmt.Sprintf("%v/projects/%v/%v/history/%v",
						globalCallbackURL, input.ProjectID, input.Script.ID, historyID),
					Success: true,
				}, input.SuccessChannel.SlackConfig.WebhookURL)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			case "webhook":
				notificationErr := notifyWebhook(ctx, historyID, input.SuccessChannel.WebhookConfig.URL, db)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			case "email":
				notificationErr := notifyEmail(ctx, historyID, input.SuccessChannel.EmailConfig.To, db)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			}
		}
	} else {
		if input.FailureChannel != nil {
			switch input.FailureChannel.Type {
			case "slack":
				notificationErr := notifySlack(ctx, SlackNotifyInput{
					ProjectName: input.ProjectName,
					ScriptName:  input.Script.Name,
					HistoryLink: fmt.Sprintf("%v/projects/%v/%v/history/%v",
						globalCallbackURL, input.ProjectID, input.Script.ID, historyID),
					Success: false,
				}, input.FailureChannel.SlackConfig.WebhookURL)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			case "webhook":
				notificationErr := notifyWebhook(ctx, historyID, input.FailureChannel.WebhookConfig.URL, db)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			case "email":
				notificationErr := notifyEmail(ctx, historyID, input.FailureChannel.EmailConfig.To, db)
				if notificationErr != nil {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &f,
						LastUsed:        &now,
					})
				} else {
					updateNotificationChannel(ctx, db, input.SuccessChannel.ID, UpdateNotificationChannelInput{
						LastUsedSuccess: &t,
						LastUsed:        &now,
					})
				}
			}
		}
	}
	return
}
