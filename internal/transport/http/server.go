package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/todo/internal/service"
)

type Server struct {
	config      Config
	taskService *service.Task
}

func New(config Config, taskService *service.Task) *Server {
	return &Server{
		config:      config,
		taskService: taskService,
	}
}

func (s *Server) Start(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	s.setRoutes(app)

	err := app.Listen(s.config.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to start http listener: %w", err)
	}

	return nil
}
