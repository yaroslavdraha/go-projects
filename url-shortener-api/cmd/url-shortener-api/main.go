package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener-api/internal/config"
)

func main() {

	config := config.MustLoad()

	fmt.Println(config)

	// todo: init config: cleanenv

	// todo: init logger: slog

	log := setupLogger(config.Env)

	log.Debug("Debug log is enabled")
	log.Info("app uses %v environment", slog.String("env", config.Env))

	// todo: init storage: sqlite

	// todo: init router: chi

	// todo: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
