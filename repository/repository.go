package repository

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

type Repository interface {
	// TODO: insert data to db
}
