package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/yekuanyshev/todo/config"
	"github.com/yekuanyshev/todo/internal/repository"
	"github.com/yekuanyshev/todo/internal/service"
	"github.com/yekuanyshev/todo/internal/transport/http"
	"github.com/yekuanyshev/todo/internal/transport/http/handler"
	"github.com/yekuanyshev/todo/pkg/logger"
	"github.com/yekuanyshev/todo/pkg/postgres"
)

func main() {
	ctx := context.Background()

	conf, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logger.New(os.Stdout, conf.LogLevel, conf.LogFormat)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	conn, err := postgres.Connect(ctx, conf.PgDSN)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	logger.Info("connected to database...")

	taskRepositoryLogger := logger.With(
		slog.String("module", "repository"),
		slog.String("domain", "task"),
	)

	taskServiceLogger := logger.With(
		slog.String("module", "service"),
		slog.String("domain", "task"),
	)

	taskHandlerLogger := logger.With(
		slog.String("module", "handler"),
		slog.String("domain", "task"),
	)

	taskRepository := repository.NewTask(conn, taskRepositoryLogger)
	taskService := service.NewTask(taskRepository, taskServiceLogger)
	taskHandler := handler.NewTask(taskService, taskHandlerLogger)

	httpConfig := http.Config{
		ListenAddress: conf.HTTPListen,
	}
	httpServer := http.New(httpConfig, taskHandler)

	logger.Info("start http server...")
	err = httpServer.Start(ctx)
	if err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}

	// todo: add graceful shutdown
}
