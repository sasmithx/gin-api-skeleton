package db

import (
	"context"
	"fmt"
	"time"

	"api-skeleton/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connectTimeout     = 5 * time.Second
	pingTimeout        = 5 * time.Second
	defaultMaxConns    = int32(10)
	defaultMinConns    = int32(1)
	defaultMaxConnIdle = 5 * time.Minute
	defaultMaxLife     = 1 * time.Hour
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) {
	if cfg.Database_Url == "" {
		return nil, fmt.Errorf("database url is required")
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.Database_Url)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	if poolConfig.MaxConns == 0 {
		poolConfig.MaxConns = defaultMaxConns
	}
	if poolConfig.MinConns == 0 {
		poolConfig.MinConns = defaultMinConns
	}
	if poolConfig.MaxConnIdleTime == 0 {
		poolConfig.MaxConnIdleTime = defaultMaxConnIdle
	}
	if poolConfig.MaxConnLifetime == 0 {
		poolConfig.MaxConnLifetime = defaultMaxLife
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), pingTimeout)
	defer pingCancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return pool, nil
}
