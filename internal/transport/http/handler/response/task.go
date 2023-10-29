package response

import (
	"time"

	"github.com/yekuanyshev/todo/internal/models"
)

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateTaskFromModel(task models.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		IsDone:    task.IsDone,
		CreatedAt: task.CreatedAt,
	}
}

func CreateTaskListFromModels(tasks []models.Task) []Task {
	var result []Task
	for _, t := range tasks {
		result = append(result, CreateTaskFromModel(t))
	}
	return result
}

type CreatedTask struct {
	ID int64 `json:"id"`
}
