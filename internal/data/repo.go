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
	DeleteOrderDetails(orderID int) error
	NewPizzaDetails(pizza *NewPizza) error
	ShowPizzaDetails(pizza *NewPizza) error
	ListAllOrdersDetails(pizza *NewPizza) error
}

// repo represents db actions object
type repo struct {
	*sql.DB
}

// NewRepo creates a new repo object.
func New(db *sql.DB) *repo {
	return &repo{db}
}

//data execution protection
