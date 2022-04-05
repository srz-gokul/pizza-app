package data

import (
	"github.com/pkg/errors"
)

const newPizzaQuery = `
INSERT INTO 
pizza(name, type)
VALUES
('$1', '$2');
		`

// GetOrderDetails get user order details from DB.
func (r *repo) NewPizzaDetails(pizza *NewPizza) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "NewPizzaDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		newPizzaQuery,
		pizza.Pizza_name,
		pizza.Pizza_type,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "NewPizzaDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "NewPizzaDetails: transaction commit failed")
	}
	return nil
}
