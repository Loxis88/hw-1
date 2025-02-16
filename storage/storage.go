package storage

import (
	"hw-1/models"
	"hw-1/storage/json_storage"
)

type OrderStorage interface {
	AddOrder(order models.Order) error
	UpdateOrder(order models.Order) error
	DeleteOrder(id uint) error
	GetOrders() []models.Order
	GetOrdersByCustomer(customerID uint, lastN int) []models.Order
	FindOrder(id uint) (*models.Order, error)
	GetExpiredOrders() []models.Order
	GetOrdersHistory(limit int) ([]models.Order, error)
	GetReturnedOrders() []models.Order
}

// Функция для создания нового хранилища
func NewJsonStorage(path string) (OrderStorage, error) {
	return json_storage.New(path)
}
