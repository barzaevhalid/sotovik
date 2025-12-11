package db

import (
	"context"

	"github.com/barzaevhalid/sotovik/internal/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg configs.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.Db.Dsn)

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil

}
