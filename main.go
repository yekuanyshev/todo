package main

import (
	"context"
	"log"
	"time"

	"github.com/yekuanyshev/todo/config"
	"github.com/yekuanyshev/todo/internal/repository"
	"github.com/yekuanyshev/todo/internal/service"
	"github.com/yekuanyshev/todo/internal/transport/http"
	"github.com/yekuanyshev/todo/pkg/postgres"
)

func main() {
	ctx := context.Background()

	conf, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := postgres.ConnectWithRetry(ctx, conf.PgDSN, 3, 5*time.Second)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	taskRepository := repository.NewTask(conn)
	taskService := service.NewTask(taskRepository)

	httpConfig := http.Config{
		ListenAddress: conf.HTTPListen,
	}
	httpServer := http.New(httpConfig, taskService)
	err = httpServer.Start(ctx)
	if err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}
