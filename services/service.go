package services

import (
	"fmt"
	"time"

	"hw-1/models"
	"hw-1/storage"
)

type OrderServiceInterface interface {
	// Основные операции с заказами
	AcceptOrder(orderID uint, customerID uint, storageDate time.Time) error
	ReturnOrderToCourier(orderID uint) error
	DeliverOrders(customerID uint, orderIDs ...uint) error
	AcceptReturns(customerID uint, orderIDs ...uint) error
	GetCustomerOrders(customerID uint, limit int, inStorageOnly bool) ([]models.Order, error)
	GetOrderHistory(limit int) ([]models.Order, error)
}

// Проверка реализации интерфейса
var _ OrderServiceInterface = (*OrderService)(nil)

type OrderService struct {
	storage storage.OrderStorage
}

func New(storage storage.OrderStorage) OrderServiceInterface {
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
		UpdatedAt:    time.Now(),
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
		order.UpdatedAt = time.Now()
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
		order.UpdatedAt = time.Now()
		if err := s.storage.UpdateOrder(*order); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetCustomerOrders(customerID uint, limit int, inStorageOnly bool) ([]models.Order, error) {
	return s.storage.GetOrdersByCustomer(customerID, limit, inStorageOnly), nil
}

func (s *OrderService) ReturnedList() {
	orders := s.storage.GetExpiredOrders()

	for _, order := range orders {
		if order.Status == models.StatusReturned {
			fmt.Println(order)
		}
	}
}
