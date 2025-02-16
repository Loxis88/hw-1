package services

import (
	"fmt"
	"time"

	"hw-1/models"
	"hw-1/storage/json_storage"
)

type OrderService struct {
	storage *json_storage.Storage
}

func New(path string) *OrderService {
	storage, err := json_storage.New(path)
	if err != nil {
		panic(err)
	}
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) AcceptOrder(orderID uint, customerID uint, StorageDate time.Time) error {
	if find, _ := s.storage.FindOrder(orderID); find != nil {
		return fmt.Errorf("order already exists") // добавить кастомные ошибки типа models.ErrOrderAlreadyExists
	}

	// Проверяем, что срок хранения не в прошлом
	if StorageDate.Before(time.Now()) {
		return fmt.Errorf("срок хранения в прошлом")
	}

	newOrder := models.Order{
		ID:          orderID,
		CustomerID:  customerID,
		StorageDate: StorageDate,
		Status:      "accepted", // доавить константы models.OrderStatusAccepted, models.OrderStatusIssued
	}

	err := s.storage.AddOrder(newOrder)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) ReturnOrder(orderID uint) error {
	order, err := s.storage.FindOrder(orderID)
	if err != nil {
		return err
	}
	if order.Status == "issued" {
		return fmt.Errorf("order already issued")
	}
	s.storage.DeleteOrder(order.ID)
	return nil
}

func (s *OrderService) Test() {
	s.storage.Test()
}
