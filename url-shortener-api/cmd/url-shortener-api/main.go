package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
	"url-shortener-api/internal/config"
	"url-shortener-api/internal/lib/logger"
	"url-shortener-api/internal/storage"
)

func main() {
	// Init config
	cnf := config.MustLoad()

	// Init logger
	log := setupLogger(cnf.Env)
	log.Debug("Debug log is enabled")
	log.Info("app uses environment", slog.String("env", cnf.Env))

	// Init storage
	db, err := storage.New(cnf.SqliteStoragePath)
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}
	_ = db

	err = db.SaveURL("https://www.google.com", "google")

	if err != nil {
		log.Error("failed to save URL", logger.Err(err))
		os.Exit(1)
	}

	// Init router: chi
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

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
