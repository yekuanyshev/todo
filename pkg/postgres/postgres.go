package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return pool, nil
}

func ConnectWithRetry(
	ctx context.Context,
	dsn string,
	attempts int,
	delay time.Duration,
) (*pgxpool.Pool, error) {
	var (
		pool *pgxpool.Pool
		err  error
	)

	for i := 0; i < attempts; i++ {
		isLast := i == attempts-1

		pool, err = Connect(ctx, dsn)
		if err != nil {
			if isLast {
				break
			}
			time.Sleep(delay)
		}
	}

	return pool, err
}
