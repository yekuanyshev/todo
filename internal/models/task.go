package models

import "time"

type Task struct {
	ID        int64
	Title     string
	IsDone    bool
	CreatedAt time.Time
}
