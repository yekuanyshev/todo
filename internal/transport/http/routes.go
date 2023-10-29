package http

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	taskRouter := app.Group("/task")
	{
		taskRouter.Get("", s.taskHandler.GetAll)
		taskRouter.Get("/:id", s.taskHandler.GetByID)
		taskRouter.Post("", s.taskHandler.Create)
		taskRouter.Put("/:id/done", s.taskHandler.Done)
		taskRouter.Put("/:id/undone", s.taskHandler.Undone)
		taskRouter.Delete("/:id", s.taskHandler.Delete)
	}
}
