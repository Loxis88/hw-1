package models

type OrderStorage interface {
	AddOrder(order Order) error
	UpdateOrder(order Order) error
	DeleteOrder(id uint) error
	GetOrders() []Order
	GetOrdersByCustomer(customerID uint, lastN int, inStorageOnly bool) []Order
	FindOrder(id uint) (*Order, error)
	GetExpiredOrders() []Order
}
