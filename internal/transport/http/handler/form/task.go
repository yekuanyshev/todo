package form

import "github.com/yekuanyshev/todo/internal/models"

type NewTask struct {
	Title string `json:"title"`
}

func (t NewTask) ToModel() models.Task {
	return models.Task{
		Title: t.Title,
	}
}
