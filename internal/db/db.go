package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(databaseURL string, maxConns, minConns int, maxIdleTime, dbTimeout string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = int32(maxConns)
	cfg.MinConns = int32(minConns)

	idleDuration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	cfg.MaxConnIdleTime = idleDuration
	cfg.MaxConnLifetime = time.Hour
	cfg.HealthCheckPeriod = time.Minute

	timeoutDuration, err := time.ParseDuration(dbTimeout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
