package main

import (
	"log"
	"log/slog"

	"github.com/Bolatyerbol/Newgoproject/internal/config"
	"github.com/Bolatyerbol/Newgoproject/internal/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	slog.Info("Error message")
	slog.Debug("Debug message")
	slog.Warn("Warn message")
	slog.Error("Error message")
}
