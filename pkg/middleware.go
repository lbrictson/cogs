package pkg

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) frontendAuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := s.sessionStore.Get(c.Request(), sessionName)
		if err != nil {
			session.Options.MaxAge = -1
			err = session.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		sessionID := session.Values["id"]
		if sessionID == nil {
			session.Options.MaxAge = -1
			err = session.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		sessionData, err := s.sessionManager.GetSession(sessionID.(string))
		if err != nil {
			session.Options.MaxAge = -1
			err = session.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		if sessionData.ExpiresAt.Before(time.Now()) {
			session.Options.MaxAge = -1
			err = session.Save(c.Request(), c.Response())
			s.sessionManager.DeleteSession(sessionID.(string))
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		c.Set("email", sessionData.Email)
		c.Set("role", sessionData.Role)
		return next(c)
	}
}

func (s *Server) globalAdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Get("role")
		if r == nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		if r.(string) != "admin" {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		return next(c)
	}
}

func (s *Server) projectAccessRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		r := c.Get("role")
		if r == nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		if r.(string) == "admin" {
			return next(c)
		}
		email := c.Get("email")
		if email == nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		projectID := c.Param("projectID")
		if projectID == "" {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// convert to int
		i, err := strconv.Atoi(projectID)
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// get the user
		user, err := getUserByEmail(ctx, s.db, email.(string))
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// get project permissions
		permissions, err := getUserAccesses(ctx, s.db, user.ID)
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// check if user has access to project
		for _, p := range permissions {
			if p.ProjectID == i {
				return next(c)
			}
		}
		return c.Render(http.StatusForbidden, "unauthorized", nil)
	}
}

func (s *Server) projectAdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		r := c.Get("role")
		if r == nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		if r.(string) == "admin" {
			return next(c)
		}
		email := c.Get("email")
		if email == nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		projectID := c.Param("projectID")
		if projectID == "" {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// convert to int
		i, err := strconv.Atoi(projectID)
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// get the user
		user, err := getUserByEmail(ctx, s.db, email.(string))
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// get project permissions
		permissions, err := getUserAccesses(ctx, s.db, user.ID)
		if err != nil {
			return c.Render(http.StatusForbidden, "unauthorized", nil)
		}
		// check if user has access to project
		for _, p := range permissions {
			if p.ProjectID == i && p.Role == "admin" {
				return next(c)
			}
		}
		return c.Render(http.StatusForbidden, "unauthorized", nil)
	}
}

func (s *Server) apiKeyAdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")
		if apiKey == "" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "unauthorized"})
		}
		valid, user := validateAPIKey(apiKey)
		if !valid {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "unauthorized"})
		}
		if user.Role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "unauthorized"})
		}
		c.Set("email", user.Email)
		c.Set("role", user.Role)
		return next(c)
	}
}
