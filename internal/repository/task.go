package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/todo/internal/models"
)

type Task struct {
	conn    *pgxpool.Pool
	builder squirrel.StatementBuilderType
}

func NewTask(conn *pgxpool.Pool) *Task {
	return &Task{
		conn:    conn,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *Task) ListAll(ctx context.Context) (result []models.Task, err error) {
	query, args, err := repo.builder.
		Select("id", "title", "is_done", "created_at").
		From("task").OrderBy("created_at DESC").ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := repo.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to run query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.IsDone, &task.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		result = append(result, task)
	}

	return result, nil
}

func (repo *Task) ByID(ctx context.Context, id int64) (result models.Task, err error) {
	query, args, err := repo.builder.
		Select("id", "title", "is_done", "created_at").
		From("task").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to build query: %w", err)
	}

	err = repo.conn.QueryRow(ctx, query, args...).Scan(
		&result.ID,
		&result.Title,
		&result.IsDone,
		&result.CreatedAt,
	)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to run query: %w", err)
	}

	return result, nil
}

func (repo *Task) Create(ctx context.Context, task models.Task) (id int64, err error) {
	query, args, err := repo.builder.
		Insert("task").
		SetMap(map[string]any{
			"title": task.Title,
		}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	err = repo.conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to run query: %w", err)
	}

	return id, nil
}

func (repo *Task) SetDone(ctx context.Context, id int64, isDone bool) (err error) {
	query, args, err := repo.builder.
		Update("task").
		Set("is_done", isDone).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}

func (repo *Task) Delete(ctx context.Context, id int64) (err error) {
	query, args, err := repo.builder.
		Delete("task").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}
