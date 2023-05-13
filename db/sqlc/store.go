package db

import "database/sql"

type Store interface {
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
	}
}
