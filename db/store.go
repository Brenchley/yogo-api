package yogo

import (
	"database/sql"
)

// Store defines all functions to execute db queries
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries
type SQLStore struct {
	yogo *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(yogo *sql.DB) Store {
	return &SQLStore{
		yogo:    yogo,
		Queries: New(yogo),
	}
}
