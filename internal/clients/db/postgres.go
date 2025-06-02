package dbclient

import (
	"context"
	"fmt"
	"videobin/internal/clients"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dsn string) (clients.DataBase, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("new.pgxpool %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool.ping %w", err)
	}

	return &psql{
		db: pool,
	}, nil
}

type psql struct {
	db *pgxpool.Pool
}

// Close implements clients.DBClient.
func (p *psql) Close() {
	p.db.Close()
}

// DB implements clients.DBClient.
func (p *psql) DB() *pgxpool.Pool {
	return p.db
}
