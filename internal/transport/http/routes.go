package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/todo/internal/transport/http/handler"
)

func (s *Server) setRoutes(app *fiber.App) {
	taskHandler := handler.NewTask(s.taskService)

	taskRouter := app.Group("/task")
	{
		taskRouter.Get("", taskHandler.GetAll)
		taskRouter.Get("/:id", taskHandler.GetByID)
		taskRouter.Post("", taskHandler.Create)
		taskRouter.Put("/:id/done", taskHandler.Done)
		taskRouter.Put("/:id/undone", taskHandler.Undone)
		taskRouter.Delete("/:id", taskHandler.Delete)
	}
}
