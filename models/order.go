package models

import (
	"errors"
	"time"
)

type OrderStatus string

const (
	StatusNew       OrderStatus = "new"
	StatusExpired   OrderStatus = "expired"
	StatusReturned  OrderStatus = "returned"
	StatusDelivered OrderStatus = "delivered"
)

var (
	ErrOrderAlreadyExists       = errors.New("заказ уже существует")            // заказ уже существует
	ErrInvalidStorageDate       = errors.New("некорректная дата хранения")      // некорректная дата хранения
	ErrOrderCannotBeReturned    = errors.New("заказ не может быть возвращен")   // заказ не может быть возвращен
	ErrOrderCannotBeDelivered   = errors.New("заказ не может быть выдан")       // заказ не может быть выдан
	ErrStoragePeriodExpired     = errors.New("срок хранения истек")             // срок хранения истек
	ErrStoragePeriodNotExpired  = errors.New("срок хранения еще не истек")      // срок хранения еще не истек
	ErrReturnPeriodExpired      = errors.New("срок возврата истек")             // срок возврата истек
	ErrOrderNotBelongToCustomer = errors.New("заказ не принадлежит покупателю") // заказ не принадлежит покупателю
)

type Order struct {
	ID           uint        `json:"id"`
	CustomerID   uint        `json:"customer_id"`
	StorageUntil time.Time   `json:"storage_until"`
	Status       OrderStatus `json:"status"`
	DeliveredAt  time.Time   `json:"delivered_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
