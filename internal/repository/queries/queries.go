package queries

import (
	_ "embed"
)

//go:embed list_all_tasks.sql
var ListAllTasks string

//go:embed get_task_by_id.sql
var GetTaskByID string

//go:embed insert_task.sql
var InsertTask string

//go:embed set_done_to_task.sql
var SetDoneToTask string

//go:embed delete_task.sql
var DeleteTask string
