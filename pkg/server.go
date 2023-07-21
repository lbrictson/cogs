package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/cogs/ent"
	"github.com/lbrictson/cogs/web"
	"github.com/robfig/cron/v3"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"strings"
)

var sessionName = "cogs-session"
var dataDirectory = ""
var globalCallbackURL = ""

type Server struct {
	port           int
	db             *ent.Client
	sessionStore   *sessions.CookieStore
	sessionManager *SessionManager
	cronService    *cron.Cron
	callbackURL    string
}

type NewServerInput struct {
	Port        int
	DB          *ent.Client
	DevMode     bool
	CallbackURL string
}

func NewServer(input NewServerInput) *Server {
	cookieSecret := strings.Replace(uuid.New().String(), "-", "", -1)
	if input.DevMode {
		fmt.Println("WARNING: Running in development mode. Cookie secret is not secure.")
		cookieSecret = "notAGreatSecretValue"
	}
	return &Server{
		port:           input.Port,
		db:             input.DB,
		sessionStore:   sessions.NewCookieStore([]byte(cookieSecret)),
		sessionManager: NewSessionManager(),
		cronService:    cron.New(),
		callbackURL:    input.CallbackURL,
	}
}

func (s *Server) Run(ctx context.Context) {
	go runErroredJobCleaner(ctx, s.db)
	e := echo.New()
	e.Renderer = mustNewRenderer()
	e.HideBanner = true
	globalCallbackURL = s.callbackURL
	// Read in static assets from the mock file system
	fSys, err := fs.Sub(web.Assets, "static")
	if err != nil {
		panic(err)
	}
	// Serve static files from virtual file system 'static' directory
	assetHandler := http.FileServer(http.FS(fSys))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	loginRequiredRoutes := e.Group("", s.frontendAuthRequired)
	// Project routes
	loginRequiredRoutes.GET("/", renderProjectsPage(ctx, s.db), s.frontendAuthRequired)
	loginRequiredRoutes.GET("/projects/create", renderCreateProjectsPage(ctx), s.globalAdminRequired)
	loginRequiredRoutes.POST("/projects/create", formCreateProject(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID", renderViewProjectPage(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.DELETE("/projects/:projectID", hookDeleteProject(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID/:script_id", renderViewScriptPage(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.POST("/projects/:projectID/:script_id", formRunScript(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.GET("/projects/:projectID/permissions", renderViewProjectPermissionsPage(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.POST("/projects/:projectID/permissions", formUpdateProjectPermissions(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID/create", renderCreateScriptPage(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.POST("/projects/:projectID/create", formCreateScript(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID/:script_id/edit", renderEditScriptPage(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.POST("/projects/:projectID/:script_id/edit", formUpdateScript(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.DELETE("/projects/:projectID/:script_id", hookDeleteScript(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID/:script_id/history", renderHistoryPage(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.GET("/projects/:projectID/:script_id/history/:historyID", renderSingleHistoryPage(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.GET("/projects/:projectID/secrets", renderViewProjectSecretsPage(ctx, s.db), s.projectAccessRequired)
	loginRequiredRoutes.GET("/projects/:projectID/secrets/create", renderViewCreateSecretPage(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.POST("/projects/:projectID/secrets/create", formCreateSecret(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.GET("/projects/:projectID/secrets/edit/:secretID", renderViewUpdateSecretPage(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.POST("/projects/:projectID/secrets/edit/:secretID", formUpdateSecret(ctx, s.db), s.projectAdminRequired)
	loginRequiredRoutes.DELETE("/projects/:projectID/secrets/delete/:secretID", hookDeleteSecret(ctx, s.db), s.projectAdminRequired)
	// Notification routes
	loginRequiredRoutes.GET("/notifications", renderNotificationsPage(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/notifications/create/:type", renderCreateNotificationPage(ctx), s.globalAdminRequired)
	loginRequiredRoutes.POST("/notifications/create/:type", formCreateNotificationChannel(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/notifications/:id", renderEditNotificationPage(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.POST("/notifications/:id", formUpdateNotificationChannel(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.DELETE("/notifications/:id", hookDeleteNotificationChannel(ctx, s.db), s.globalAdminRequired)
	// User routes
	loginRequiredRoutes.GET("/users", renderUsersPage(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/users/create", renderCreateUsersPage(ctx), s.globalAdminRequired)
	loginRequiredRoutes.POST("/users/create", formCreateUser(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.GET("/users/:id", renderEditUsersPage(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.POST("/users/:id", formEditUser(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.DELETE("/users/:id", hookDeleteUser(ctx, s.db), s.globalAdminRequired)
	// Login routes
	e.GET("/login", renderLoginPage(ctx))
	e.POST("/login", formLogin(ctx, s.db, s.sessionStore, s.sessionManager))
	e.GET("/failed_login", renderFailedLoginPage(ctx))
	e.GET("/logout", logoutHook(ctx, s.sessionStore, s.sessionManager))
	startScheduledJobs(ctx, s.db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.port)))
}

func mustNewRenderer() *Renderer {
	r := Renderer{}
	r.templates = template.Must(template.ParseFS(web.Assets, "templates/*.tmpl"))
	return &r
}

type Renderer struct {
	templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	templateData := make(map[string]interface{})
	if data != nil {
		casted, ok := data.(map[string]interface{})
		if ok {
			templateData = casted
		} else {
			LogFromCtx(c.Request().Context()).Error("Failed to cast template data")
			return errors.New("failed to cast template data")
		}
	}
	templateData["Role"] = c.Get("role")
	templateData["Email"] = c.Get("email")
	err := t.templates.ExecuteTemplate(w, name, templateData)
	if err != nil {
		LogFromCtx(c.Request().Context()).Error(err.Error())
		return err
	}
	return nil
}
