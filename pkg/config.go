package pkg

import (
	"flag"
	"github.com/lbrictson/cogs/ent"
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
		Location: *dataDir,
	})
	if err != nil {
		panic(err)
	}
	dataDirectory = *dataDir
	return &Config{
		DataDirectory: *dataDir,
		Port:          *port,
		DBConnection:  dbConn,
		LogLevel:      *level,
		LogFormat:     *format,
		DevMode:       *devMode,
	}
}
