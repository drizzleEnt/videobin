package clients

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase interface {
	DB() *pgxpool.Pool
	Close()
}
