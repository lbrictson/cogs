package pkg

import (
	"flag"
	"github.com/lbrictson/cogs/ent"
	"os"
	"strconv"
)

type Config struct {
	DataDirectory string
	Port          int
	DBConnection  *ent.Client
	LogLevel      string
	LogFormat     string
	DevMode       bool
}

func NewConfig() *Config {
	dataDir := flag.String("data", "tmp", "location of data directory")
	port := flag.Int("port", 8080, "port to listen on")
	level := flag.String("level", "info", "log level")
	format := flag.String("format", "text", "log format (text or json)")
	devMode := flag.Bool("dev", false, "development mode")
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
		DataDirectory: checkEnvVarForStringValue("COGS_DATA", *dataDir),
		Port:          checkEnvVarForIntValue("COGS_PORT", *port),
		DBConnection:  dbConn,
		LogLevel:      checkEnvVarForStringValue("COGS_LOG_LEVEL", *level),
		LogFormat:     checkEnvVarForStringValue("COGS_LOG_FORMAT", *format),
		DevMode:       checkEnvVarForBoolValue("COGS_DEV_MODE", *devMode),
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
