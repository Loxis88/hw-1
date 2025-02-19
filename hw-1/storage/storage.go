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
	FindOrder(id uint) (*models.Order, error)
}

func NewJsonStorage(path string) (OrderStorage, error) {
	return json_storage.New(path)
}
