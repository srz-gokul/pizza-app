package data

// MockDB provides a static data backed implementation of Repo.
type MockDB struct {
	DBStatusInfoFn         func() (string, error)
	AddOrderDetailsFn      func(order *OrderData) error
	GetOrderDetailsFn      func(id int) (*OrderData, error)
	UpdateOrderDetailsFn   func(order *OrderData) error
	DeleteOrderDetailsFn   func(id int) error
	NewPizzaDetailsFn      func(pizza *NewPizza) error
	ShowPizzaDetailsFn     func(pizza *NewPizza) error
	ListAllOrdersDetailsFn func(pizza *NewPizza) error
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

func (db MockDB) DeleteOrderDetails(id int) error {
	return db.DeleteOrderDetailsFn(id)
}

func (db MockDB) NewPizzaDetails(pizza *NewPizza) error {
	return db.NewPizzaDetailsFn(pizza)
}

func (db MockDB) ShowPizzaDetails(pizza *NewPizza) error {
	return db.ShowPizzaDetailsFn(pizza)
}

func (db MockDB) ListAllOrdersDetails(pizza *NewPizza) error {
	return db.ListAllOrdersDetailsFn(pizza)
}
