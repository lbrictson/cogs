package pkg

import (
	"context"
	"github.com/lbrictson/cogs/ent"
	"github.com/robfig/cron/v3"
	"sync"
)

var cronLock sync.Mutex
var cronTracker map[int]cron.EntryID
var cronInstance *cron.Cron

func startScheduledJobs(ctx context.Context, db *ent.Client) {
	scripts, err := getScripts(ctx, db)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	for _, script := range scripts {
		if script.ScheduleCron != "" {
			if script.ScheduleEnabled {
				err := upsertScheduledJobToScheduler(ctx, script.ID, script.ScheduleCron, db)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
				} else {
					LogFromCtx(ctx).With("schedule", script.ScheduleCron).With("script", script.Name).Info("Scheduled job added")
				}
			}
		}
	}
}

func checkCron() {
	if cronInstance == nil {
		cronInstance = cron.New()
		cronInstance.Start()
	}
	if cronTracker == nil {
		cronTracker = make(map[int]cron.EntryID)
	}
}

func removeScheduledJobFromScheduler(scriptID int) {
	checkCron()
	cronLock.Lock()
	defer cronLock.Unlock()
	// Check if the job is already in the tracker map
	_, ok := cronTracker[scriptID]
	if ok {
		// If it is, remove it from the cron scheduler
		cronInstance.Remove(cronTracker[scriptID])
		// remove form map
		delete(cronTracker, scriptID)
	}
	return
}

func upsertScheduledJobToScheduler(ctx context.Context, scriptID int, schedule string, db *ent.Client) error {
	checkCron()
	// Remove it from the scheduler if it's already there
	removeScheduledJobFromScheduler(scriptID)
	cronLock.Lock()
	defer cronLock.Unlock()
	// Add it to the scheduler
	id, err := cronInstance.AddFunc(schedule, func() {
		// Run the script
		runScheduledJob(ctx, db, scriptID)
	})
	if err != nil {
		return err
	}
	// Add it to the tracker map
	cronTracker[scriptID] = id
	return nil
}

func runScheduledJob(ctx context.Context, db *ent.Client, scriptID int)  {
	// Get the script
	script, err := getScriptByID(ctx, db, scriptID)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	project, err := getProjectByID(ctx, db, script.ProjectID)
	if err != nil {
		LogFromCtx(ctx).Error(err.Error())
		return
	}
	args := make(map[string]string)
	// Run the script
	runnerInput := RunScriptInput{
		 Script:         script,
		 Caller:         "system",
		 Trigger:        "scheduled",
		 Args:           args,
		 ProjectID:      script.ProjectID,
		 ProjectName:    project.Name,
		 SuccessChannel: nil,
		 FailureChannel: nil,
	 }
	if script.SuccessNotificationID != nil {
		if *script.SuccessNotificationID != 0 {
			sNotify, err := getNotificationChannelByID(ctx, db, *script.SuccessNotificationID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return
			}
			runnerInput.SuccessChannel = &sNotify
		}
	}
	if script.FailureNotificationID != nil {
		if *script.FailureNotificationID != 0 {
			fNotify, err := getNotificationChannelByID(ctx, db, *script.FailureNotificationID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return
			}
			runnerInput.FailureChannel = &fNotify
		}
	}
	jid := runScript(ctx, db, runnerInput)
	LogFromCtx(ctx).With("jid", jid).With("script", script.Name).Info("Scheduled job started")
}