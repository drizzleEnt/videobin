package postgres

import (
	"videobin/internal/clients"
	"videobin/internal/repository"
)

func NewDatabaseStorage(db clients.DataBase) repository.DatabaseStorage {
	return &repo{
		db: db,
	}
}

type repo struct {
	db clients.DataBase
}
