package data

import "github.com/pkg/errors"

const addOrderDetailsQuery = `
	INSERT INTO pizza_order
	(
		pizza_id,
		pizza_size,
		cooking_stage,
		user_id,
		start_time,
		is_active
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	);
`

// AddOrderDetails is used to add pizza order details into the DB.
func (r *repo) AddOrderDetails(order *OrderData) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "AddOrderDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		addOrderDetailsQuery,
		order.PizzaID,
		order.PizzaSize,
		order.CookingStage,
		order.UserID,
		order.StartTime,
		order.IsActive,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "AddOrderDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "AddOrderDetails: transaction commit failed")
	}
	return nil
}

const getUserDetailsQuery = `
		SELECT
			id,
			pizza_id,
			pizza_size,
			cooking_stage,
			user_id,
			start_time
		FROM   
			pizza_order
		WHERE  
			user_id = $1 AND 
			is_active = true
		;
		`

// GetOrderDetails get user order details from DB.
func (r *repo) GetOrderDetails(id int) (*OrderData, error) {
	var order OrderData
	if err := r.QueryRow(getUserDetailsQuery, id).Scan(
		&order.ID,
		&order.PizzaID,
		&order.PizzaSize,
		&order.CookingStage,
		&order.UserID,
		&order.StartTime,
	); err != nil {
		return nil, errors.Wrap(err, "GetOrderDetails: row scan failed")
	}
	return &order, nil
}

const updateOrderDetailsQuery = `
	UPDATE 
		pizza_order
	SET
		cooking_stage = $1,
		end_time 	  = $2,
		is_active     = $3
	WHERE
		id = $4
	;
`

// UpdateOrderDetails update order details in DB.
func (r *repo) UpdateOrderDetails(order *OrderData) error {
	tx, err := r.Begin()
	if err != nil {
		return errors.Wrap(err, "UpdateOrderDetails: transaction begin failed")
	}
	_, err = tx.Exec(
		updateOrderDetailsQuery,
		order.CookingStage,
		order.EndTime,
		order.IsActive,
		order.ID,
	)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "UpdateOrderDetails: transaction execution failed")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "UpdateOrderDetails: transaction commit failed")
	}
	return nil
}
