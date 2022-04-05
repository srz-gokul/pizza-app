package data

import "github.com/pkg/errors"

const deleteOrderDetailsQuery = `
		DELETE *
		FROM   
			pizza_order
		WHERE  
			id = $1 AND 
			is_active = true
		;
		`

// GetOrderDetails get user order details from DB.
func (r *repo) DeleteOrderDetails(orderID int) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "DeleteOrderDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		deleteOrderDetailsQuery,
		orderID,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "DeleteOrderDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "DeleteOrderDetails: transaction commit failed")
	}
	return nil
}
