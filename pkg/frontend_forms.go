package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/cogs/ent"
	"github.com/lbrictson/cogs/ent/schema"
	"github.com/robfig/cron/v3"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func formCreateProject(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FormData struct {
			Name        string `form:"name"`
			Description string `form:"description"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			return err
		}
		_, err := createProject(ctx, db, NewProjectInput{
			Name:        data.Name,
			Description: data.Description,
		})
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		LogFromCtx(ctx).Info("created project", "name", data.Name, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/")
	}
}

func formCreateUser(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FormData struct {
			Email    string `form:"email"`
			Password string `form:"password"`
			Role     string `form:"role"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			return err
		}
		// validate password
		isPasswordValid, reason := validatePassword(data.Password)
		if !isPasswordValid {
			return c.Render(http.StatusBadRequest, "generic_error", map[string]interface{}{
				"Message": reason,
			})
		}
		_, err := createUser(ctx, db, CreateUserInput{
			Email:    data.Email,
			Password: data.Password,
			Role:     data.Role,
			APIKey:   generateAPIKey(),
		})
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		LogFromCtx(ctx).Info("created user", "email", data.Email, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/users")
	}
}

func formEditUser(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		type FormData struct {
			Email    string `form:"email"`
			Password string `form:"password"`
			Role     string `form:"role"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		input := UpdateUserInput{}
		if data.Password != "" {
			input.Password = &data.Password
			// validate password
			isPasswordValid, reason := validatePassword(data.Password)
			if !isPasswordValid {
				return c.Render(http.StatusBadRequest, "generic_error", map[string]interface{}{
					"Message": reason,
				})
			}
		}
		if data.Role != "" {
			input.Role = &data.Role
		}
		_, err = updateUser(ctx, db, i, input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		LogFromCtx(ctx).Info("edited user", "name", data.Email, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/users")
	}
}

func formLogin(ctx context.Context, db *ent.Client, store *sessions.CookieStore, sessionMgr *SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FormData struct {
			Email    string `form:"email"`
			Password string `form:"password"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			return c.Redirect(http.StatusFound, "/failed_login")
		}
		data.Email = strings.ToLower(data.Email)
		u, err := getUserByEmail(ctx, db, data.Email)
		if err != nil {
			LogFromCtx(ctx).Warn(err.Error())
			return c.Redirect(http.StatusFound, "/failed_login")
		}
		if !comparePasswordHashes(u.HashedPassword, data.Password) {
			LogFromCtx(ctx).Warn("invalid password for user " + data.Email)
			return c.Redirect(http.StatusFound, "/failed_login")
		}
		session, err := store.Get(c.Request(), sessionName)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		session.Values["id"] = sessionMgr.CreateSession(SessionData{
			Email:     u.Email,
			Role:      u.Role,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 1), // 1 day
		})
		session.Options.MaxAge = 60 * 60 * 24 * 1 // 1 day
		session.Options.HttpOnly = true
		session.Options.Secure = false
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return err
		}
		LogFromCtx(ctx).Info("successful login", "user", data.Email)
		return c.Redirect(http.StatusFound, "/")
	}
}

func formUpdateProjectPermissions(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("projectID")
		i, err := strconv.Atoi(id)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		formData, err := c.FormParams()
		if err != nil {
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		projectPermissions, err := getProjectAccesses(ctx, db, i)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		for k, v := range formData {
			if len(v) > 0 {
				if k != "" {
					userID, err := strconv.Atoi(k)
					if err != nil {
						return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
							"Message": err.Error(),
						})
					}
					// We need to update this users permission, if they already have access we might need to change it
					// if they don't have access we need to add it
					for _, p := range projectPermissions {
						if p.ProjectID == i && p.UserID == userID {
							// This user already has access
							// remove it so we can add their new access
							deleteAccess(ctx, db, p.ID)
						}
					}
					// Add the new access
					if v[0] == "admin" || v[0] == "user" {
						LogFromCtx(ctx).Info("adding access " + v[0] + " for user " + k + " to project " + id)
						input := CreateAccessInput{
							ProjectID: i,
							UserID:    userID,
							Role:      v[0],
						}
						_, err := createAccess(ctx, db, input)
						if err != nil {
							return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
								"Message": err.Error(),
							})
						}
					}
				}
			}
		}
		LogFromCtx(ctx).Info("updated project permissions", "project", id, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+id)
	}
}

func formUpdateScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		scriptID := c.Param("script_id")
		// convert to int
		i, err := strconv.Atoi(scriptID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		_, err = getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		type FormData struct {
			Name                  string `form:"name"`
			Description           string `form:"description"`
			Timeout               int    `form:"timeout"`
			Script                string `form:"script"`
			Params                string `form:"params"`
			ScheduleEnabled       string `form:"scheduleEnabled"`
			Schedule              string `form:"schedule"`
			ActualScheduleEnabled bool
			SuccessNotification   int `form:"successNotification"`
			FailureNotification   int `form:"failureNotification"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if data.ScheduleEnabled == "on" {
			data.ActualScheduleEnabled = true
			// Need to validate schedule
			_, err := cron.ParseStandard(data.Schedule)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": fmt.Sprintf("Failed to validate cron schedule %v: %v", data.Schedule, err.Error()),
				})
			}
		}
		updatedOptionsData := []schema.ScriptInputOptions{}
		// Validate param data
		updatedOptionsNeeded := true
		options := []schema.ScriptInputOptions{}
		if data.Params != "" {
			err = json.Unmarshal([]byte(data.Params), &options)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			for _, o := range options {
				// Make sure name does not contain spaces
				if strings.Contains(o.Name, " ") {
					LogFromCtx(ctx).Error("Param name cannot contain spaces")
					return errors.New("option name cannot contain spaces")
				}
			}
			updatedOptionsData = options
		}
		data.Script = strings.ReplaceAll(data.Script, "\r\n", "\n")
		data.Script = strings.TrimSpace(data.Script)
		input := UpdateScriptInput{
			Name:                         &data.Name,
			Script:                       &data.Script,
			Description:                  &data.Description,
			TimeoutSeconds:               &data.Timeout,
			ScheduleCron:                 &data.Schedule,
			ScheduleEnabled:              &data.ActualScheduleEnabled,
			SuccessNotificationChannelID: &data.SuccessNotification,
			FailureNotificationChannelID: &data.FailureNotification,
		}
		if updatedOptionsNeeded {
			input.Parameters = &updatedOptionsData
		}
		_, err = updateScript(ctx, db, i, input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if *input.ScheduleEnabled {
			upsertScheduledJobToScheduler(ctx, i, *input.ScheduleCron, db)
		} else {
			removeScheduledJobFromScheduler(i)
		}
		LogFromCtx(ctx).Info("updated script", "script", scriptID, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+c.Param("projectID")+"/"+scriptID)
	}
}

func formCreateScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("projectID")
		// convert to int
		i, err := strconv.Atoi(projectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		type FormData struct {
			Name                  string `form:"name"`
			Description           string `form:"description"`
			Timeout               int    `form:"timeout"`
			Script                string `form:"script"`
			Params                string `form:"params"`
			ScheduleEnabled       string `form:"scheduleEnabled"`
			Schedule              string `form:"schedule"`
			ActualScheduleEnabled bool
			SuccessNotification   int `form:"successNotification"`
			FailureNotification   int `form:"failureNotification"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if data.ScheduleEnabled == "on" {
			data.ActualScheduleEnabled = true
			// Need to validate schedule
			_, err := cron.ParseStandard(data.Schedule)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": fmt.Sprintf("Failed to validate cron schedule %v: %v", data.Schedule, err.Error()),
				})
			}
		}
		updatedOptionsData := []schema.ScriptInputOptions{}
		// Validate param data
		options := []schema.ScriptInputOptions{}
		if data.Params != "" {
			err = json.Unmarshal([]byte(data.Params), &options)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			for _, o := range options {
				// Make sure name does not contain spaces
				if strings.Contains(o.Name, " ") {
					LogFromCtx(ctx).Error("Param name cannot contain spaces")
					return errors.New("option name cannot contain spaces")
				}
			}
			updatedOptionsData = options
		}
		data.Script = strings.ReplaceAll(data.Script, "\r\n", "\n")
		data.Script = strings.TrimSpace(data.Script)
		input := CreateScriptInput{
			Name:                         data.Name,
			Script:                       data.Script,
			Description:                  data.Description,
			TimeoutSeconds:               data.Timeout,
			ProjectID:                    i,
			Parameters:                   updatedOptionsData,
			ScheduleCron:                 data.Schedule,
			ScheduleEnabled:              data.ActualScheduleEnabled,
			SuccessNotificationChannelID: &data.SuccessNotification,
			FailureNotificationChannelID: &data.FailureNotification,
		}
		script, err := createScript(ctx, db, input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if input.ScheduleEnabled {
			upsertScheduledJobToScheduler(ctx, script.ID, input.ScheduleCron, db)
		}
		LogFromCtx(ctx).Info("created script", "script", data.Name, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+c.Param("projectID")+"/"+fmt.Sprintf("%v", script.ID))
	}
}

func formRunScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		scriptID := c.Param("script_id")
		// convert to int
		i, err := strconv.Atoi(scriptID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		s, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		args := make(map[string]string)
		for _, x := range s.Parameters {
			args[x.Name] = c.FormValue(x.Name)
		}
		// Get the project
		p, err := getProjectByID(ctx, db, s.ProjectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		runnerInput := RunScriptInput{
			Script:      s,
			Caller:      c.Get("email").(string),
			Trigger:     "webUI",
			Args:        args,
			ProjectID:   s.ProjectID,
			ProjectName: p.Name,
		}
		if s.SuccessNotificationID != nil {
			if *s.SuccessNotificationID != 0 {
				sNotify, err := getNotificationChannelByID(ctx, db, *s.SuccessNotificationID)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
					return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
						"Message": err.Error(),
					})
				}
				runnerInput.SuccessChannel = &sNotify
			}
		}
		if s.FailureNotificationID != nil {
			if *s.FailureNotificationID != 0 {
				fNotify, err := getNotificationChannelByID(ctx, db, *s.FailureNotificationID)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
					return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
						"Message": err.Error(),
					})
				}
				runnerInput.FailureChannel = &fNotify
			}
		}
		runScript(ctx, db, runnerInput)
		LogFromCtx(ctx).Info("ran script", "script", scriptID, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+c.Param("projectID")+"/"+scriptID+"/history")
	})
}

func formCreateSecret(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		id := c.Param("projectID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		project, err := getProjectByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		type FormData struct {
			Name  string `form:"name"`
			Value string `form:"value"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		_, err = createSecret(ctx, db, CreateSecretInput{
			Name:      data.Name,
			Value:     data.Value,
			ProjectID: project.ID,
			CreatedBy: userFromEchoContext(c),
		})
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		LogFromCtx(ctx).Info("created secret", "secret", data.Name, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+c.Param("projectID")+"/secrets")
	})
}

func formUpdateSecret(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		id := c.Param("secretID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		secret, err := getSecretByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		type FormData struct {
			Value string `form:"value"`
		}
		var data FormData
		if err := c.Bind(&data); err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		_, err = updateSecret(ctx, db, secret.ID, UpdateSecretInput{
			Value: &data.Value,
		})
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		LogFromCtx(ctx).Info("updated secret", "secret", secret.Name, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/projects/"+c.Param("projectID")+"/secrets")
	})
}

func formCreateNotificationChannel(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		switch c.Param("type") {
		case "slack":
			type FormData struct {
				Name string `form:"name"`
				URL  string `form:"url"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err := createNotificationChannel(ctx, db, CreateNotificationChannelInput{
				Name:        data.Name,
				Type:        "slack",
				SlackConfig: schema.SlackConfig{WebhookURL: data.URL},
				Enabled:     true,
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			LogFromCtx(ctx).Info("created notification channel", "name", data.Name, "user", userFromEchoContext(c))
			return c.Redirect(http.StatusFound, "/notifications")
		case "email":
			type FormData struct {
				Name  string `form:"name"`
				Email string `form:"email"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err := createNotificationChannel(ctx, db, CreateNotificationChannelInput{
				Name:        data.Name,
				Type:        "email",
				EmailConfig: schema.EmailConfig{To: data.Email},
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			LogFromCtx(ctx).Info("created notification channel", "name", data.Name, "user", userFromEchoContext(c))
			return c.Redirect(http.StatusFound, "/notifications")
		case "webhook":
			type FormData struct {
				Name string `form:"name"`
				URL  string `form:"url"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err := createNotificationChannel(ctx, db, CreateNotificationChannelInput{
				Name:          data.Name,
				Type:          "webhook",
				WebhookConfig: schema.WebhookConfig{URL: data.URL},
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			LogFromCtx(ctx).Info("created notification channel", "name", data.Name, "user", userFromEchoContext(c))
			return c.Redirect(http.StatusFound, "/notifications")
		default:
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": "unknown notification channel type",
			})
		}
	})
}

func formUpdateNotificationChannel(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		id := c.Param("id")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		notificationChannel, err := getNotificationChannelByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		switch notificationChannel.Type {
		case "slack":
			type FormData struct {
				Name string `form:"name"`
				URL  string `form:"url"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err = updateNotificationChannel(ctx, db, notificationChannel.ID, UpdateNotificationChannelInput{
				Name:        &data.Name,
				SlackConfig: &schema.SlackConfig{WebhookURL: data.URL},
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
		case "email":
			type FormData struct {
				Name  string `form:"name"`
				Email string `form:"email"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err = updateNotificationChannel(ctx, db, notificationChannel.ID, UpdateNotificationChannelInput{
				Name:        &data.Name,
				EmailConfig: &schema.EmailConfig{To: data.Email},
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
		case "webhook":
			type FormData struct {
				Name string `form:"name"`
				URL  string `form:"url"`
			}
			var data FormData
			if err := c.Bind(&data); err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			_, err = updateNotificationChannel(ctx, db, notificationChannel.ID, UpdateNotificationChannelInput{
				Name:          &data.Name,
				WebhookConfig: &schema.WebhookConfig{URL: data.URL},
			})
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
		default:
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": "unknown notification channel type",
			})
		}
		LogFromCtx(ctx).Info("updated notification channel", "name", notificationChannel.Name, "user", userFromEchoContext(c))
		return c.Redirect(http.StatusFound, "/notifications")
	})
}
