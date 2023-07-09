package pkg

import (
	"github.com/dustin/go-humanize"
	"github.com/lbrictson/cogs/ent/schema"
	"time"
)

type UserModel struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	APIKey         string    `json:"api_key"`
	Role           string    `json:"role"`
	HashedPassword string    `json:"-"`
}

type ProjectModel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
}

type ScriptModel struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Description           string    `json:"description"`
	ProjectID             int       `json:"project_id"`
	TimeoutSeconds        int       `json:"timeout_seconds"`
	Parameters            []schema.ScriptInputOptions
	Script                string `json:"script"`
	SuccessNotificationID *int   `json:"success_notification_id"`
	FailureNotificationID *int   `json:"failure_notification_id"`
	ScheduleEnabled       bool   `json:"schedule_enabled"`
	ScheduleCron          string `json:"schedule_cron"`
}

type AccessModel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ProjectID int       `json:"project_id"`
	UserID    int       `json:"user_id"`
	Role      string    `json:"role"`
}

type SecretModel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
}

type HistoryModel struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ScriptID    int       `json:"script_id"`
	Trigger     string    `json:"trigger"`
	Success     bool      `json:"success"`
	TriggeredBy string    `json:"triggered_by"`
	ExitCode    int       `json:"exit_code"`
	Duration    int       `json:"duration"`
	Output      string    `json:"output"`
	RunID       string    `json:"run_id"`
	Status      string    `json:"status"`
	Arguments   map[string]string
}

func (s HistoryModel) HumanizeCreatedAt() string {
	return humanize.Time(s.CreatedAt)
}

func (s HistoryModel) FormatTime() string {
	return s.CreatedAt.Format("2006-01-02@15:04")
}

type FrontendScriptBlobModel struct {
	Parameters []schema.ScriptInputOptions `json:"parameters"`
}

type ScriptStatsModel struct {
	ID              int       `json:"id"`
	ScriptID        int       `json:"script_id"`
	ProjectID       int       `json:"project_id"`
	TotalRuns       int       `json:"total_runs"`
	TotalSuccess    int       `json:"total_success"`
	TotalError      int       `json:"total_error"`
	AverageDuration int       `json:"average_duration"`
	LastRun         time.Time `json:"last_run"`
	LongestRun      int       `json:"longest_run"`
	ShortestRun     int       `json:"shortest_run"`
	TotalDuration   int       `json:"total_duration"`
	SuccessRate     float64   `json:"success_rate"`
}

func (s *ScriptStatsModel) HumanizeLastRun() string {
	return humanize.Time(s.LastRun)
}

func (s *ScriptStatsModel) FormatTimeLastRun() string {
	return s.LastRun.Format("2006-01-02@15:04")
}

type NotificationChannelModel struct {
	ID            int                  `json:"id"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Name          string               `json:"name"`
	Type          string               `json:"type"`
	SlackConfig   schema.SlackConfig   `json:"slack_config"`
	EmailConfig   schema.EmailConfig   `json:"email_config"`
	WebhookConfig schema.WebhookConfig `json:"webhook_config"`
	Enabled       bool                 `json:"enabled"`
}
