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
	return &OrderService{storage: storage}
}

// Принять заказ от курьера
func (s *OrderService) AcceptOrder(orderID uint, customerID uint, storageDate time.Time) error {
	if find, _ := s.storage.FindOrder(orderID); find != nil {
		return models.ErrOrderAlreadyExists
	}

	if storageDate.Before(time.Now()) {
		return models.ErrInvalidStorageDate
	}

	newOrder := models.Order{
		ID:           orderID,
		CustomerID:   customerID,
		StorageUntil: storageDate,
		Status:       models.StatusNew,
	}

	return s.storage.AddOrder(newOrder)
}

// Вернуть заказ курьеру
func (s *OrderService) ReturnOrderToCourier(orderID uint) error {
	order, err := s.storage.FindOrder(orderID)
	if err != nil {
		return err
	}

	if order.Status == models.StatusNew {
		return models.ErrOrderCannotBeReturned
	}

	if time.Now().Before(order.StorageUntil) {
		return models.ErrStoragePeriodNotExpired
	}

	return s.storage.DeleteOrder(orderID)
}

// Выдать заказы
func (s *OrderService) DeliverOrders(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return err
		}

		if order.CustomerID != customerID {
			return fmt.Errorf("%w : %d", models.ErrOrderNotBelongToCustomer, id)
		}

		if order.Status != models.StatusNew {
			return models.ErrOrderCannotBeDelivered
		}

		order.Status = models.StatusDelivered
		if err := s.storage.UpdateOrder(*order); err != nil {
			return err
		}
	}

	return nil
}

// принять возвраты клиента
func (s *OrderService) AcceptReturns(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return err
		}

		if order.CustomerID != customerID {
			return fmt.Errorf("%w : %d", models.ErrOrderNotBelongToCustomer, id)
		}

		if order.Status != models.StatusDelivered {
			return models.ErrOrderCannotBeReturned
		}

		order.Status = models.StatusReturned
		if err := s.storage.UpdateOrder(*order); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetCustomerOrders(customerID uint, limit int, inStorageOnly bool) ([]models.Order, error) {
	return s.storage.GetOrdersByCustomer(customerID, limit, inStorageOnly), nil
}
