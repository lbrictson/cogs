package pkg

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/cogs/ent"
	"github.com/lbrictson/cogs/ent/schema"
	"github.com/robfig/cron/v3"
	"net/http"
	"strconv"
)

func apiV1GetProjects(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, err := getProjects(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, projects, "  ")
	}
}

func apiV1GetScripts(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("projectID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		scripts, err := getProjectScripts(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, scripts, "  ")
	}
}

func apiV1GetProjectScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("scriptID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, script, "  ")
	}
}

func apiV1GetScriptHistories(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("scriptID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		limit := 50
		offset := 0
		if c.QueryParam("limit") != "" {
			limit, err = strconv.Atoi(c.QueryParam("limit"))
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		if c.QueryParam("offset") != "" {
			offset, err = strconv.Atoi(c.QueryParam("offset"))
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		history, err := getScriptHistories(ctx, db, i, limit, offset)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, history, "  ")
	}
}

func apiV1GetScriptHistory(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("historyID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		history, err := getHistoryByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, history, "  ")
	}
}

func apiV1RunScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("scriptID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		type Input struct {
			Arguments map[string]string `json:"arguments"`
		}
		var input Input
		err = c.Bind(&input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		// Make sure all arguments are provided
		for _, arg := range script.Parameters {
			found := false
			for k := range input.Arguments {
				if k == arg.Name {
					found = true
				}
			}
			if !found {
				return c.JSON(http.StatusBadRequest, fmt.Sprintf("Argument '%s' is required", arg.Name))
			}
		}
		project, err := getProjectByID(ctx, db, script.ProjectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		runInput := RunScriptInput{
			Script:         script,
			Caller:         c.Get("email").(string),
			Trigger:        "API",
			Args:           input.Arguments,
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
					return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
						"Message": err.Error(),
					})
				}
				runInput.SuccessChannel = &sNotify
			}
		}
		if script.FailureNotificationID != nil {
			if *script.FailureNotificationID != 0 {
				fNotify, err := getNotificationChannelByID(ctx, db, *script.FailureNotificationID)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
					return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
						"Message": err.Error(),
					})
				}
				runInput.FailureChannel = &fNotify
			}
		}
		history := runScript(ctx, db, runInput)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		type RunResponse struct {
			HistoryID int `json:"historyID"`
		}
		return c.JSONPretty(http.StatusOK, RunResponse{HistoryID: history}, "  ")
	}
}

func apiV1UpdateScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("scriptID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		type Input struct {
			Name                  *string `json:"name"`
			Description           *string `json:"description"`
			TimeoutSeconds        *int    `json:"timeout_seconds"`
			Parameters            *[]schema.ScriptInputOptions
			Script                *string `json:"script"`
			SuccessNotificationID *int    `json:"success_notification_id"`
			FailureNotificationID *int    `json:"failure_notification_id"`
			ScheduleEnabled       *bool   `json:"schedule_enabled"`
			ScheduleCron          *string `json:"schedule_cron"`
		}
		var input Input
		err = c.Bind(&input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		updaterInput := UpdateScriptInput{}
		if input.Name != nil {
			// Make sure name is not empty
			if *input.Name == "" {
				return c.JSON(http.StatusBadRequest, "Name cannot be empty")
			}
			updaterInput.Name = input.Name
		}
		if input.Description != nil {
			updaterInput.Description = input.Description
		}
		if input.TimeoutSeconds != nil {
			// Make sure timeout is greater than 1
			if *input.TimeoutSeconds < 1 {
				return c.JSON(http.StatusBadRequest, "Timeout must be greater than 1")
			}
			updaterInput.TimeoutSeconds = input.TimeoutSeconds
		}
		if input.Parameters != nil {
			updaterInput.Parameters = input.Parameters
		}
		if input.Script != nil {
			// Make sure script is not empty
			if *input.Script == "" {
				return c.JSON(http.StatusBadRequest, "Script cannot be empty")
			}
			updaterInput.Script = input.Script
		}
		if input.SuccessNotificationID != nil {
			// Make sure the notification channel exists
			_, err := getNotificationChannelByID(ctx, db, *input.SuccessNotificationID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			updaterInput.SuccessNotificationChannelID = input.SuccessNotificationID
		}
		if input.FailureNotificationID != nil {
			// Make sure the notification channel exists
			_, err := getNotificationChannelByID(ctx, db, *input.FailureNotificationID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			updaterInput.FailureNotificationChannelID = input.FailureNotificationID
		}
		if input.ScheduleEnabled != nil {
			updaterInput.ScheduleEnabled = input.ScheduleEnabled
		}
		if input.ScheduleCron != nil {
			// Validate the cron schedule
			_, err := cron.ParseStandard(*input.ScheduleCron)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			updaterInput.ScheduleCron = input.ScheduleCron
		}
		updatedScript, err := updateScript(ctx, db, script.ID, updaterInput)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONPretty(http.StatusOK, updatedScript, "  ")
	}
}
