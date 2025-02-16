package api

import "hw-1/models"

// OrderService определяет интерфейс для работы с заказами на уровне бизнес-логики
type OrderService interface {
	CreateOrder(ID, customerID uint) (uint, error)
	GetOrders() []models.Order
	GetCustomerOrders(customerID uint, lastN int, inStorageOnly bool) []models.Order
	TakeOrder(orderID uint) error
	CancelOrder(orderID uint) error
	ProcessExpiredOrders() error
}
