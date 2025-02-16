package models

import (
	"time"
)

type Order struct {
	ID          uint      `json:"id"`
	CustomerID  uint      `json:"customer_id"`
	StorageDate time.Time `json:"storage_date"`
	Status      string    `json:"status"`
}

type Customer struct {
	CustomerID uint    `json:"customer_id"`
	Orders     []Order `json:"orders"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id uint) (*Order, error)
	Update(order *Order) error
	Delete(id uint) error
}
