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
		err = deleteUser(ctx, db, i)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			return c.HTML(http.StatusInternalServerError, err.Error())
		}
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
