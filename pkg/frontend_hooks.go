package pkg

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/cogs/ent"
	"net/http"
	"strconv"
)

func hookDeleteProject(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		err = deleteProject(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.HTML(http.StatusInternalServerError, err.Error())
		}
		// Delete all scripts in the project
		s, _ := getProjectScripts(ctx, db, i)
		for _, script := range s {
			err = deleteScript(ctx, db, script.ID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.HTML(http.StatusInternalServerError, err.Error())
			}
			removeScheduledJobFromScheduler(script.ID)
		}
		// Delete all secrets in the project
		secrets, _ := getProjectSecrets(ctx, db, i)
		for _, secret := range secrets {
			err = deleteSecret(ctx, db, secret.ID)
			if err != nil {
				LogFromCtx(ctx).Error(err.Error())
				return c.HTML(http.StatusInternalServerError, err.Error())
			}
		}
		LogFromCtx(ctx).Info("Deleted project", "id", i)
		return c.HTML(http.StatusOK, "<h1>Deleted</h1>")
	}
}

func hookDeleteUser(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		userToDelete, err := getUserByID(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		err = deleteUser(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		removeAPIKeyFromCache(userToDelete.APIKey)
		LogFromCtx(ctx).Info("Deleted user", "id", i)
		return c.HTML(http.StatusOK, "<h1>Deleted</h1>")
	}
}

func hookDeleteSecret(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("secretID")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		err = deleteSecret(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.HTML(http.StatusInternalServerError, err.Error())
		}
		LogFromCtx(ctx).Info("Deleted secret", "id", i)
		return c.HTML(http.StatusOK, "<h1>Deleted</h1>")
	}
}

func hookDeleteScript(ctx context.Context, db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("script_id")
		// convert to int
		i, err := strconv.Atoi(id)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		err = deleteScript(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.HTML(http.StatusInternalServerError, err.Error())
		}
		removeScheduledJobFromScheduler(i)
		LogFromCtx(ctx).Info("Deleted script", "id", i)
		return c.HTML(http.StatusOK, "<h1>Deleted</h1>")
	}
}

func hookDeleteNotificationChannel(ctx context.Context, db *ent.Client) echo.HandlerFunc {
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
		err = deleteNotificationChannel(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		// Remove the notification channel from any scripts that use it
		allScripts, err := getScripts(ctx, db)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
				"Message": err.Error(),
			})
		}
		none := int(0)
		for _, script := range allScripts {
			updateNeeded := false
			updateContent := UpdateScriptInput{
				SuccessNotificationChannelID: script.SuccessNotificationID,
				FailureNotificationChannelID: script.FailureNotificationID,
			}
			if script.SuccessNotificationID != nil {
				if *script.SuccessNotificationID == i {
					updateContent.SuccessNotificationChannelID = &none
					updateNeeded = true
				}
			}
			if script.FailureNotificationID != nil {
				if *script.FailureNotificationID == i {
					updateContent.FailureNotificationChannelID = &none
					updateNeeded = true
				}
			}
			if updateNeeded {
				_, err = updateScript(ctx, db, script.ID, updateContent)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
					return c.Render(http.StatusInternalServerError, "generic_error", map[string]interface{}{
						"Message": err.Error(),
					})
				}
			}
		}
		LogFromCtx(ctx).Info("Deleted notification channel", "id", i)
		return c.HTML(http.StatusOK, "<h1>Deleted</h1>")
	}
}

func logoutHook(ctx context.Context, cookieStore *sessions.CookieStore, sessionMgr *SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := cookieStore.Get(c.Request(), sessionName)
		if err != nil {
			return c.Redirect(http.StatusFound, "/")
		}
		sessionID := session.Values["id"]
		if sessionID == nil {
			return c.Redirect(http.StatusFound, "/")
		}
		sessionMgr.DeleteSession(sessionID.(string))
		session.Options.MaxAge = -1
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.Redirect(http.StatusFound, "/")
		}
		return c.Redirect(http.StatusFound, "/")
	}
}
