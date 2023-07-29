package pkg

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/cogs/ent"
	"net/http"
	"strconv"
	"strings"
)

func renderProjectsPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, err := getProjects(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "projects", map[string]interface{}{
			"Projects": projects,
		})
	}
}

func renderCreateProjectsPage(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "create_projects", nil)
	}
}

func renderViewProjectPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		scriptsForProject, err := getProjectScripts(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "project", map[string]interface{}{
			"Project": project,
			"Scripts": scriptsForProject,
		})
	}
}

func renderViewProjectSecretsPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		secrets, err := getProjectSecrets(ctx, db, project.ID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "project_secrets", map[string]interface{}{
			"Project": project,
			"Secrets": secrets,
		})
	}
}

func renderViewCreateSecretPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		return c.Render(http.StatusOK, "create_secret", map[string]interface{}{
			"Project": project,
		})
	}
}

func renderViewUpdateSecretPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		sID := c.Param("secretID")
		// convert to int
		sI, err := strconv.Atoi(sID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		secret, err := getSecretByID(ctx, db, sI)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "edit_secret", map[string]interface{}{
			"Project": project,
			"Secret":  secret,
		})
	}
}

func renderViewProjectPermissionsPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		permissions, err := getProjectAccesses(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		users, err := getUsers(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		type PermissionMapping struct {
			Email         string
			IsAdmin       bool
			IsUser        bool
			HasNoAccess   bool
			IsGlobalAdmin bool
			UserID        int
		}
		permissionMappings := []PermissionMapping{}
		for _, user := range users {
			found := false
			for _, permission := range permissions {
				if user.ID == permission.UserID {
					found = true
					m := PermissionMapping{
						Email:         user.Email,
						IsAdmin:       false,
						IsUser:        false,
						HasNoAccess:   true,
						IsGlobalAdmin: false,
						UserID:        user.ID,
					}
					if permission.Role == "admin" {
						m.IsAdmin = true
						m.HasNoAccess = false
					}
					if permission.Role == "user" {
						m.IsUser = true
						m.HasNoAccess = false
					}
					permissionMappings = append(permissionMappings, m)
				}
			}
			if !found {
				m := PermissionMapping{
					Email:         user.Email,
					IsAdmin:       false,
					IsUser:        false,
					HasNoAccess:   true,
					IsGlobalAdmin: false,
					UserID:        user.ID,
				}
				if user.Role == "admin" {
					m.IsGlobalAdmin = true
					m.HasNoAccess = false
				}
				permissionMappings = append(permissionMappings, m)
			}
		}
		return c.Render(http.StatusOK, "permissions", map[string]interface{}{
			"Project":     project,
			"Permissions": permissionMappings,
			"Users":       users,
		})
	}
}

func renderViewScriptPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		projectID := c.Param("projectID")
		// convert to int
		j, err := strconv.Atoi(projectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		project, err := getProjectByID(ctx, db, j)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		lines := strings.Split(script.Script, "\n")
		numberOfLines := len(lines)
		return c.Render(http.StatusOK, "script", map[string]interface{}{
			"Script":  script,
			"Project": project,
			"Lines":   numberOfLines,
		})
	}
}

func renderCreateScriptPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		project, err := getProjectByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		notificationChannelsAvailable, err := getNotificationChannels(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": "Failed to get notification channels",
			})
		}
		return c.Render(http.StatusOK, "create_script", map[string]interface{}{
			"Project":  project,
			"Channels": notificationChannelsAvailable,
		})
	}
}

func renderEditScriptPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		projectID := c.Param("projectID")
		// convert to int
		j, err := strconv.Atoi(projectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		project, err := getProjectByID(ctx, db, j)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		o, _ := json.MarshalIndent(script.Parameters, "", "  ")
		prettyParams := string(o)
		prettyParams = strings.TrimSpace(prettyParams)
		lines := strings.Split(script.Script, "\n")
		// Pretty print options
		optionLines := 0
		if script.Parameters != nil {
			j, _ := json.MarshalIndent(script.Parameters, "", "  ")
			prettyParams = string(j)
			prettyParams = strings.TrimSpace(prettyParams)
			optionLines = len(strings.Split(prettyParams, "\n"))
		}
		numberOfLines := len(lines)
		notificationChannelsAvailable, err := getNotificationChannels(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": "Failed to get notification channels",
			})
		}
		selectedFailureNotificationChannel := 0
		selectedSuccessNotificationChannel := 0
		if script.FailureNotificationID != nil {
			selectedFailureNotificationChannel = *script.FailureNotificationID
		}
		if script.SuccessNotificationID != nil {
			selectedSuccessNotificationChannel = *script.SuccessNotificationID
		}
		return c.Render(http.StatusOK, "edit_script", map[string]interface{}{
			"Script":                             script,
			"ScriptParams":                       prettyParams,
			"Project":                            project,
			"Lines":                              numberOfLines,
			"OptionsLines":                       optionLines,
			"Channels":                           notificationChannelsAvailable,
			"SelectedSuccessNotificationChannel": selectedSuccessNotificationChannel,
			"SelectedFailureNotificationChannel": selectedFailureNotificationChannel,
		})
	}
}

func renderSingleHistoryPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		historyID := c.Param("historyID")
		// convert to int
		i, err := strconv.Atoi(historyID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		history, err := getHistoryByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		script, err := getScriptByID(ctx, db, history.ScriptID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		project, err := getProjectByID(ctx, db, script.ProjectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "run_history", map[string]interface{}{
			"History": history,
			"Script":  script,
			"Project": project,
		})
	}
}

func renderHistoryPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		pageNumber := 0
		page := c.QueryParam("page")
		// convert to int
		if page != "" {
			p, err := strconv.Atoi(page)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
					"Message": err.Error(),
				})
			}
			pageNumber = p
		}
		offset := 0
		limit := 10
		if pageNumber > 1 {
			offset = (pageNumber - 1) * limit
		}
		more := false
		scriptID := c.Param("script_id")
		// convert to int
		i, err := strconv.Atoi(scriptID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		history, err := getScriptHistories(ctx, db, i, limit, offset)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if len(history) == limit {
			more = true
		}
		script, err := getScriptByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		projectID := c.Param("projectID")
		// convert to int
		j, err := strconv.Atoi(projectID)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		project, err := getProjectByID(ctx, db, j)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		stats, err := getScriptStatsByScriptID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		feModels := []FrontendScriptHistoryModel{}
		for _, h := range history {
			feModels = append(feModels, FrontendScriptHistoryModel{
				ScriptID:     h.ScriptID,
				ProjectID:    project.ID,
				ScriptName:   script.Name,
				HistoryModel: h,
			})
		}
		return c.Render(http.StatusOK, "script_history", map[string]interface{}{
			"Script":       script,
			"Project":      project,
			"History":      feModels,
			"Page":         pageNumber,
			"More":         more,
			"NextPage":     pageNumber + 1,
			"PreviousPage": pageNumber - 1,
			"Stats":        stats,
		})
	}
}

func renderUsersPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := getUsers(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		if c.QueryParams().Get("search") != "" {
			old := users
			searchString := strings.ToLower(c.QueryParams().Get("search"))
			users = []UserModel{}
			for _, user := range old {
				if strings.Contains(strings.ToLower(user.Email), searchString) {
					users = append(users, user)
				}
			}
		}
		return c.Render(http.StatusOK, "users", map[string]interface{}{
			"Users":  users,
			"Search": c.QueryParams().Get("search"),
		})
	}
}

func renderCreateUsersPage(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "create_users", nil)
	}
}

func renderEditUsersPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		user, err := getUserByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "edit_users", map[string]interface{}{
			"User": user,
		})
	}
}

func renderNotificationsPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		notifications, err := getNotificationChannels(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{"Message": err.Error()})
		}
		return c.Render(http.StatusOK, "notifications", map[string]interface{}{
			"Notifications": notifications,
		})
	}
}

func renderCreateNotificationPage(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		notificationType := c.Param("type")
		switch notificationType {
		case "slack":
			return c.Render(http.StatusOK, "create_slack_notification", nil)
		case "email":
			return c.Render(http.StatusOK, "create_email_notification", nil)
		case "webhook":
			return c.Render(http.StatusOK, "create_webhook_notification", nil)
		default:
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{"Message": "Invalid notification type"})
		}
	}
}

func renderEditNotificationPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{"Message": err.Error()})
		}
		n, err := getNotificationChannelByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{"Message": err.Error()})
		}
		switch n.Type {
		case "slack":
			return c.Render(http.StatusOK, "edit_slack_notification", map[string]interface{}{
				"Channel": n,
			})
		case "email":
			return c.Render(http.StatusOK, "edit_email_notification", map[string]interface{}{
				"Channel": n,
			})
		case "webhook":
			return c.Render(http.StatusOK, "edit_webhook_notification", map[string]interface{}{
				"Channel": n,
			})
		default:
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{"Message": "Invalid notification type"})
		}
	}
}

func renderLoginPage(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	}
}

func renderFailedLoginPage(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "failed_login", nil)
	}
}

func renderAPIKeyPage(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := getUserByEmail(ctx, db, c.Get("email").(string))
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		return c.Render(http.StatusOK, "api_key",
			map[string]interface{}{
				"APIKey": u.APIKey,
			})
	}
}
