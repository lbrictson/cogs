package pkg

import (
	"context"
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

type SMPTSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var smtpSettings SMPTSettings

type Server struct {
	port           int
	db             *ent.Client
	sessionStore   *sessions.CookieStore
	sessionManager *SessionManager
	cronService    *cron.Cron
	callbackURL    string
	retentionDays  int
	brand          string
}

type NewServerInput struct {
	Port          int
	DB            *ent.Client
	DevMode       bool
	CallbackURL   string
	RetentionDays int
	Brand         string
	SMPTHost      string
	SMPTPort      int
	SMPTUsername  string
	SMPTPassword  string
	SMPTFrom      string
}

func NewServer(input NewServerInput) *Server {
	cookieSecret := strings.Replace(uuid.New().String(), "-", "", -1)
	if input.DevMode {
		fmt.Println("WARNING: Running in development mode. Cookie secret is not secure.")
		cookieSecret = "notAGreatSecretValue"
	}
	if input.Brand == "" {
		input.Brand = "Cogs"
	}
	smtpSettings = SMPTSettings{
		Host:     input.SMPTHost,
		Port:     input.SMPTPort,
		Username: input.SMPTUsername,
		Password: input.SMPTPassword,
		From:     input.SMPTFrom,
	}
	return &Server{
		port:           input.Port,
		db:             input.DB,
		sessionStore:   sessions.NewCookieStore([]byte(cookieSecret)),
		sessionManager: NewSessionManager(),
		cronService:    cron.New(),
		callbackURL:    input.CallbackURL,
		retentionDays:  input.RetentionDays,
		brand:          input.Brand,
	}
}

func (s *Server) Run(ctx context.Context) {
	go runErroredJobCleaner(ctx, s.db)
	go runHistoryRetention(ctx, s.db, s.retentionDays)
	err := seedAPIKeyCache(ctx, s.db)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Renderer = mustNewRenderer(s.brand)
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
	// Hooks for HTMX
	loginRequiredRoutes.GET("/projects/:projectID/:script_id/history/:historyID/htmx", hookRefreshScriptHistoryFrontendContent(ctx, s.db), s.projectAccessRequired)
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
	// API key routes
	loginRequiredRoutes.GET("/api_key", renderAPIKeyPage(ctx, s.db), s.globalAdminRequired)
	loginRequiredRoutes.POST("/api_key", formRegenerateAPIKey(ctx, s.db), s.globalAdminRequired)
	// Login routes
	e.GET("/login", renderLoginPage(ctx))
	e.POST("/login", formLogin(ctx, s.db, s.sessionStore, s.sessionManager))
	e.GET("/failed_login", renderFailedLoginPage(ctx))
	e.GET("/logout", logoutHook(ctx, s.sessionStore, s.sessionManager))
	// API v1
	apiV1Routes := e.Group("/api/v1", s.apiKeyAdminRequired)
	apiV1Routes.GET("/project", apiV1GetProjects(ctx, s.db))
	apiV1Routes.GET("/project/:projectID", apiV1GetScripts(ctx, s.db))
	apiV1Routes.GET("/script/:scriptID", apiV1GetProjectScript(ctx, s.db))
	apiV1Routes.PUT("/script/:scriptID", apiV1UpdateScript(ctx, s.db))
	apiV1Routes.POST("/run/:scriptID", apiV1RunScript(ctx, s.db))
	apiV1Routes.GET("/history/:scriptID", apiV1GetScriptHistories(ctx, s.db))
	apiV1Routes.GET("/history/:scriptID/:historyID", apiV1GetScriptHistory(ctx, s.db))
	startScheduledJobs(ctx, s.db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.port)))
}

func mustNewRenderer(brand string) *Renderer {
	r := Renderer{}
	r.templates = template.Must(template.ParseFS(web.Assets, "templates/*.tmpl"))
	r.brand = brand
	return &r
}

type Renderer struct {
	templates *template.Template
	brand     string
}

func (t *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	templateData := make(map[string]interface{})
	if data != nil {
		casted, ok := data.(map[string]interface{})
		if ok {
			templateData = casted
		} else {
			renderErr := t.templates.ExecuteTemplate(w, name, data)
			if renderErr != nil {
				LogFromCtx(c.Request().Context()).Error(renderErr.Error())
				return renderErr
			}
			return nil
		}
	}
	templateData["Role"] = c.Get("role")
	templateData["Email"] = c.Get("email")
	templateData["Brand"] = t.brand
	err := t.templates.ExecuteTemplate(w, name, templateData)
	if err != nil {
		LogFromCtx(c.Request().Context()).Error(err.Error())
		return err
	}
	return nil
}
