package data

import (
	"github.com/pkg/errors"
)

const showPizzaQuery = `
SELECT * 
FROM PIZZA;
		`

// GetOrderDetails get user order details from DB.
func (r *repo) ShowPizzaDetails(pizza *NewPizza) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "ShowPizzaDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		showPizzaQuery,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ShowPizzaDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ShowPizzaDetails: transaction commit failed")
	}
	return nil
}

const ListAllOrdersDetailsQuery = `
	SELECT
	(
		pizza_id,
		pizza_size,
		user_id,
		starting_date,
		ending_date
	)
	FROM   
			pizza_details
	WHERE  
	DATE_TIME_COLUMN BETWEEN 
	'starting_date' AND 'ending_date'
	`

// ListAllorderDetails for the last days from DB.
func (r *repo) ListAllOrdersDetails(pizza *NewPizza) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "ListALLOrdersDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		ListAllOrdersDetailsQuery,
		pizza.Pizza_id,
		pizza.Pizza_size,
		pizza.user_id,
		pizza.starting_date,
		pizza.ending_date,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ListALLOrdersDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ListALLOrdersDetails: transaction commit failed")
	}
	return nil
}
