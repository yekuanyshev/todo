package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/todo/internal/service"
	"github.com/yekuanyshev/todo/internal/transport/http/handler/form"
	"github.com/yekuanyshev/todo/internal/transport/http/handler/response"
)

type Task struct {
	taskService *service.Task
	logger      *slog.Logger
}

func NewTask(taskService *service.Task, logger *slog.Logger) *Task {
	return &Task{
		taskService: taskService,
		logger:      logger,
	}
}

func (h *Task) GetAll(c *fiber.Ctx) error {
	tasks, err := h.taskService.GetAll(c.Context())
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	result := response.CreateTaskListFromModels(tasks)
	return c.JSON(response.NewSuccess(result))
}

func (h *Task) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	task, err := h.taskService.GetByID(c.Context(), int64(id))
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	result := response.CreateTaskFromModel(task)
	return c.JSON(response.NewSuccess(result))
}

func (h *Task) Create(c *fiber.Ctx) error {
	var newTask form.NewTask

	err := c.BodyParser(&newTask)
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	id, err := h.taskService.Create(c.Context(), newTask.ToModel())
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	return c.JSON(response.NewSuccess(response.CreatedTask{ID: id}))
}

func (h *Task) Done(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	err = h.taskService.Done(c.Context(), int64(id))
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	return c.JSON(response.NewSuccess(nil))
}

func (h *Task) Undone(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	err = h.taskService.Undone(c.Context(), int64(id))
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	return c.JSON(response.NewSuccess(nil))
}

func (h *Task) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	err = h.taskService.Delete(c.Context(), int64(id))
	if err != nil {
		return c.JSON(response.NewError(err))
	}

	return c.JSON(response.NewSuccess(nil))
}
