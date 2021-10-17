package data

// MockDB provides a static data backed implementation of Repo.
type MockDB struct {
	DBStatusInfoFn       func() (string, error)
	AddOrderDetailsFn    func(order *OrderData) error
	GetOrderDetailsFn    func(id int) (*OrderData, error)
	UpdateOrderDetailsFn func(order *OrderData) error
}

func (db MockDB) DBStatusInfo() (string, error) {
	return db.DBStatusInfoFn()
}

func (db MockDB) AddOrderDetails(order *OrderData) error {
	return db.AddOrderDetailsFn(order)
}

func (db MockDB) GetOrderDetails(id int) (*OrderData, error) {
	return db.GetOrderDetailsFn(id)
}

func (db MockDB) UpdateOrderDetails(order *OrderData) error {
	return db.UpdateOrderDetailsFn(order)
}
