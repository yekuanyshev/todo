package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/todo/internal/transport/http/handler"
)

type Server struct {
	config      Config
	taskHandler *handler.Task
}

func New(config Config, taskHandler *handler.Task) *Server {
	return &Server{
		config:      config,
		taskHandler: taskHandler,
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
