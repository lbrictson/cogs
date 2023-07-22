package pkg

import (
	"flag"
	"github.com/lbrictson/cogs/ent"
	"os"
	"strconv"
)

type Config struct {
	DataDirectory        string
	Port                 int
	DBConnection         *ent.Client
	LogLevel             string
	LogFormat            string
	DevMode              bool
	CallbackURL          string
	HistoryRetentionDays int
	Brand                string
	SMTPHost             string
	SMTPPort             int
	SMTPUsername         string
	SMTPPassword         string
	SMTPFrom             string
}

func NewConfig() *Config {
	dataDir := flag.String("data", "tmp", "location of data directory")
	port := flag.Int("port", 8080, "port to listen on")
	level := flag.String("level", "info", "log level")
	format := flag.String("format", "text", "log format (text or json)")
	devMode := flag.Bool("dev", false, "development mode")
	callback := flag.String("callback", "http://localhost:8080", "callback url")
	retention := flag.Int("retention", 30, "history retention in days")
	brand := flag.String("brand", "Cogs", "brand name")
	smtpHost := flag.String("smtp-host", "", "smtp host")
	smtpPort := flag.Int("smtp-port", 25, "smtp port")
	smtpUsername := flag.String("smtp-username", "", "smtp username")
	smtpPassword := flag.String("smtp-password", "", "smtp password")
	smtpFrom := flag.String("smtp-from", "cogs@localhost.com", "smtp from address")
	flag.Parse()
	dbConn, err := NewDatabaseConnection(NewDatabaseConnectionInput{
		InMemory: false,
		Location: checkEnvVarForStringValue("COGS_DATA", *dataDir),
	})
	if err != nil {
		panic(err)
	}
	dataDirectory = *dataDir
	return &Config{
		DataDirectory:        checkEnvVarForStringValue("COGS_DATA", *dataDir),
		Port:                 checkEnvVarForIntValue("COGS_PORT", *port),
		DBConnection:         dbConn,
		LogLevel:             checkEnvVarForStringValue("COGS_LOG_LEVEL", *level),
		LogFormat:            checkEnvVarForStringValue("COGS_LOG_FORMAT", *format),
		DevMode:              checkEnvVarForBoolValue("COGS_DEV_MODE", *devMode),
		CallbackURL:          checkEnvVarForStringValue("COGS_CALLBACK_URL", *callback),
		HistoryRetentionDays: checkEnvVarForIntValue("COGS_HISTORY_RETENTION_DAYS", *retention),
		Brand:                checkEnvVarForStringValue("COGS_BRAND", *brand),
		SMTPHost:             checkEnvVarForStringValue("COGS_SMTP_HOST", *smtpHost),
		SMTPPort:             checkEnvVarForIntValue("COGS_SMTP_PORT", *smtpPort),
		SMTPUsername:         checkEnvVarForStringValue("COGS_SMTP_USERNAME", *smtpUsername),
		SMTPPassword:         checkEnvVarForStringValue("COGS_SMTP_PASSWORD", *smtpPassword),
		SMTPFrom:             checkEnvVarForStringValue("COGS_SMTP_FROM", *smtpFrom),
	}
}

func checkEnvVarForStringValue(envVar string, defaultValue string) string {
	e := os.Getenv(envVar)
	if e != "" {
		return e
	}
	return defaultValue
}

func checkEnvVarForIntValue(envVar string, defaultValue int) int {
	e := os.Getenv(envVar)
	if e != "" {
		i, castErr := strconv.Atoi(e)
		if castErr != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}

func checkEnvVarForBoolValue(envVar string, defaultValue bool) bool {
	e := os.Getenv(envVar)
	if e != "" {
		b, castErr := strconv.ParseBool(e)
		if castErr != nil {
			return defaultValue
		}
		return b
	}
	return defaultValue
}
