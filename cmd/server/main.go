package main

import (
	"context"
	"github.com/lbrictson/cogs/pkg"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

func main() {
	c := pkg.NewConfig()
	level := slog.LevelInfo
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	if strings.ToLower(c.LogFormat) == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	}
	// Validate the scripts folder exists
	if _, err := os.Stat(c.DataDirectory + "/scripts"); os.IsNotExist(err) {
		// Create the directory
		err := os.MkdirAll(c.DataDirectory+"/scripts", 0755)
		if err != nil {
			panic(err)
		}
	}
	ctx := pkg.LogIntoCtx(context.Background(), logger)
	seedErr := pkg.ExecuteDefaultSeeds(ctx, c.DBConnection)
	if seedErr != nil {
		panic(seedErr)
	}
	logger.Info("starting cogs server")
	s := pkg.NewServer(pkg.NewServerInput{
		Port:          c.Port,
		DB:            c.DBConnection,
		DevMode:       c.DevMode,
		CallbackURL:   c.CallbackURL,
		RetentionDays: c.HistoryRetentionDays,
		Brand:         c.Brand,
		SMPTHost:      c.SMTPHost,
		SMPTPort:      c.SMTPPort,
		SMPTUsername:  c.SMTPUsername,
		SMPTPassword:  c.SMTPPassword,
		SMPTFrom:      c.SMTPFrom,
	})
	s.Run(ctx)
}
