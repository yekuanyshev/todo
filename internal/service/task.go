package service

import (
	"context"
	"log/slog"

	"github.com/yekuanyshev/todo/internal/models"
	"github.com/yekuanyshev/todo/internal/repository"
)

type Task struct {
	taskRepository *repository.Task
	logger         *slog.Logger
}

func NewTask(taskRepository *repository.Task, logger *slog.Logger) *Task {
	return &Task{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (srv *Task) GetAll(ctx context.Context) (result []models.Task, err error) {
	return srv.taskRepository.ListAll(ctx)
}

func (srv *Task) GetByID(ctx context.Context, id int64) (result models.Task, err error) {
	return srv.taskRepository.ByID(ctx, id)
}

func (srv *Task) Create(ctx context.Context, task models.Task) (id int64, err error) {
	return srv.taskRepository.Create(ctx, task)
}

func (srv *Task) Done(ctx context.Context, id int64) (err error) {
	return srv.taskRepository.SetDone(ctx, id, true)
}

func (srv *Task) Undone(ctx context.Context, id int64) (err error) {
	return srv.taskRepository.SetDone(ctx, id, false)
}

func (srv *Task) Delete(ctx context.Context, id int64) (err error) {
	return srv.taskRepository.Delete(ctx, id)
}
