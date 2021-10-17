package data

import (
	"database/sql"
)

// Repo outlines the data access layer.
type Repo interface {
	DBStatusInfo() (string, error)
	AddOrderDetails(order *OrderData) error
	GetOrderDetails(id int) (*OrderData, error)
	UpdateOrderDetails(order *OrderData) error
}

// repo represents db actions object
type repo struct {
	*sql.DB
}

// NewRepo creates a new repo object.
func New(db *sql.DB) *repo {
	return &repo{db}
}
