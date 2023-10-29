package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/todo/internal/models"
)

type Task struct {
	conn    *pgxpool.Pool
	logger  *slog.Logger
	builder squirrel.StatementBuilderType
}

func NewTask(conn *pgxpool.Pool, logger *slog.Logger) *Task {
	return &Task{
		conn:    conn,
		logger:  logger,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *Task) ListAll(ctx context.Context) (result []models.Task, err error) {
	query, args := repo.builder.
		Select("id", "title", "is_done", "created_at").
		From("task").OrderBy("created_at DESC").MustSql()

	logger := repo.logger.With(
		slog.String("func", "ListAll"),
		slog.String("query", query),
		slog.Any("args", args),
	)

	err = pgxscan.Select(ctx, repo.conn, &result, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return result, nil
}

func (repo *Task) ByID(ctx context.Context, id int64) (result models.Task, err error) {
	query, args := repo.builder.
		Select("id", "title", "is_done", "created_at").
		From("task").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		MustSql()

	logger := repo.logger.With(
		slog.String("func", "ByID"),
		slog.Int64("id", id),
		slog.String("query", query),
		slog.Any("args", args),
	)

	err = pgxscan.Get(ctx, repo.conn, &result, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return result, nil
}

func (repo *Task) Create(ctx context.Context, task models.Task) (id int64, err error) {
	query, args := repo.builder.
		Insert("task").
		SetMap(map[string]any{
			"title": task.Title,
		}).
		Suffix("RETURNING id").
		MustSql()

	logger := repo.logger.With(
		slog.String("func", "Create"),
		slog.Any("task", task),
		slog.String("query", query),
		slog.Any("args", args),
	)

	err = pgxscan.Get(ctx, repo.conn, &id, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to run query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return id, nil
}

func (repo *Task) SetDone(ctx context.Context, id int64, isDone bool) (err error) {
	query, args := repo.builder.
		Update("task").
		Set("is_done", isDone).
		Where(squirrel.Eq{"id": id}).
		MustSql()

	logger := repo.logger.With(
		slog.String("func", "SetDone"),
		slog.Int64("id", id),
		slog.String("query", query),
		slog.Any("args", args),
	)

	_, err = repo.conn.Exec(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to exec query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return nil
}

func (repo *Task) Delete(ctx context.Context, id int64) (err error) {
	query, args := repo.builder.
		Delete("task").
		Where(squirrel.Eq{"id": id}).
		MustSql()

	logger := repo.logger.With(
		slog.String("func", "Delete"),
		slog.Int64("id", id),
		slog.String("query", query),
		slog.Any("args", args),
	)

	_, err = repo.conn.Exec(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to exec query: %w", err)
		logger.Error("error", slog.Any("err", err))
		return
	}

	logger.Debug("success")
	return nil
}
