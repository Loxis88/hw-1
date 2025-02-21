package services

import (
	"hw-1/models"
	"time"
)

// Принять заказ от курьера
func (s *OrderService) AcceptOrder(orderID uint, customerID uint, storageDate time.Time) error {
	order, _ := s.storage.FindOrder(orderID)
	if order != nil {
		return ErrOrderAlreadyExists
	}

	// валидация времени хранения заказа
	if storageDate.Before(time.Now()) {
		return ErrInvalidStorageDate
	}

	newOrder := models.Order{
		ID:           orderID,
		CustomerID:   customerID,
		StorageUntil: storageDate,
		Status:       models.StatusNew,
		UpdatedAt:    time.Now(),
	}

	return s.storage.AddOrder(newOrder)
}
