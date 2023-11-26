package repository

import (
	"context"
	"fmt"
	"github.com/yekuanyshev/todo/internal/repository/queries"
	"log/slog"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/todo/internal/models"
)

type Task struct {
	conn   *pgxpool.Pool
	logger *slog.Logger
}

func NewTask(conn *pgxpool.Pool, logger *slog.Logger) *Task {
	return &Task{
		conn:   conn,
		logger: logger,
	}
}

func (repo *Task) ListAll(ctx context.Context) (result []models.Task, err error) {
	query := queries.ListAllTasks
	logger := repo.logger.With(
		slog.String("func", "ListAll"),
		slog.String("query", query),
	)

	err = pgxscan.Select(ctx, repo.conn, &result, query)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return result, nil
}

func (repo *Task) ByID(ctx context.Context, id int64) (result models.Task, err error) {
	query := queries.GetTaskByID
	logger := repo.logger.With(
		slog.String("func", "ByID"),
		slog.Int64("id", id),
		slog.String("query", query),
	)

	err = pgxscan.Get(ctx, repo.conn, &result, query, id)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return result, nil
}

func (repo *Task) Create(ctx context.Context, task models.Task) (id int64, err error) {
	query := queries.InsertTask
	logger := repo.logger.With(
		slog.String("func", "Create"),
		slog.Any("task", task),
		slog.String("query", query),
	)

	err = pgxscan.Get(ctx, repo.conn, &id, query, task.Title)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return id, nil
}

func (repo *Task) SetDone(ctx context.Context, id int64, isDone bool) (err error) {
	query := queries.SetDoneToTask
	logger := repo.logger.With(
		slog.String("func", "SetDone"),
		slog.Int64("id", id),
		slog.String("query", query),
	)

	_, err = repo.conn.Exec(ctx, query, isDone, id)
	if err != nil {
		err = fmt.Errorf("failed to exec query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return nil
}

func (repo *Task) Delete(ctx context.Context, id int64) (err error) {
	query := queries.DeleteTask
	logger := repo.logger.With(
		slog.String("func", "Delete"),
		slog.Int64("id", id),
		slog.String("query", query),
	)

	_, err = repo.conn.Exec(ctx, query, id)
	if err != nil {
		err = fmt.Errorf("failed to exec query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return nil
}
