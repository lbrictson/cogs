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
}

func NewConfig() *Config {
	dataDir := flag.String("data", "tmp", "location of data directory")
	port := flag.Int("port", 8080, "port to listen on")
	level := flag.String("level", "info", "log level")
	format := flag.String("format", "text", "log format (text or json)")
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
	}
}
