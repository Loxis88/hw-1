package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"hw-1/models"
	"hw-1/storage"
)

type OrderServiceInterface interface {
	AcceptOrder(orderID uint, customerID uint, storageDate time.Time) error
	ReturnOrderToCourier(orderID uint) error
	DeliverOrders(customerID uint, orderIDs ...uint) error
	AcceptReturns(customerID uint, orderIDs ...uint) error
	GetCustomerOrders(customerID uint, limit int) ([]models.Order, error)
	GetOrderHistory(limit int) ([]models.Order, error)
	GetReturnedOrders(page, pageSize int) ([]models.Order, error)
	ImportOrders(path string) error
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
	order, _ := s.storage.FindOrder(orderID)
	if order != nil {
		return models.ErrOrderAlreadyExists
	}

	// валидация времени хранения заказа
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

func (s *OrderService) ReturnOrderToCourier(orderID uint) error {
	order, err := s.storage.FindOrder(orderID)
	if err != nil {
		return err
	}

	if order.Status == models.StatusNew || order.Status == models.StatusDelivered {
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

		if err := s.isOrderBelongsToCustomer(order, customerID); err != nil {
			return err
		}

		if order.Status != models.StatusNew {
			return models.ErrOrderCannotBeDelivered
		}

		order.Status = models.StatusDelivered
		order.UpdatedAt, order.DeliveredAt = time.Now(), time.Now()
		if err := s.storage.UpdateOrder(*order); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) isOrderBelongsToCustomer(order *models.Order, customerID uint) error {
	if order.CustomerID != customerID {
		return fmt.Errorf("%w : %d", models.ErrOrderNotBelongToCustomer, order.ID)
	}
	return nil
}

func (s *OrderService) canOrderBeReturned(order *models.Order) error {
	if order.Status != models.StatusDelivered || time.Now().Sub(order.DeliveredAt) > 48*time.Hour {
		return models.ErrOrderCannotBeReturned
	}
	return nil
}

func (s *OrderService) canOrderBeDelivered(order *models.Order) error {
	if order.Status != models.StatusNew {
		return models.ErrOrderCannotBeDelivered
	}
	return nil
}

func (s *OrderService) updateOrderStatus(order *models.Order, status models.OrderStatus) error {
	order.Status = status
	order.UpdatedAt = time.Now()
	return s.storage.UpdateOrder(*order)
}

func (s *OrderService) deleteOrder(orderID uint) error {
	return s.storage.DeleteOrder(orderID)
}

// принять возвраты клиента
func (s *OrderService) AcceptReturns(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return err
		}

		if err := s.isOrderBelongsToCustomer(order, customerID); err != nil {
			return err
		}

		if err := s.canOrderBeReturned(order); err != nil {
			return err
		}
	}
	for _, id := range orderIDs {
		order, _ := s.storage.FindOrder(id)
		if err := s.updateOrderStatus(order, models.StatusReturned); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetCustomerOrders(customerID uint, limit int) ([]models.Order, error) {
	orders := s.storage.GetOrders()
	customerOrders := []models.Order{}

	if len(orders) == 0 {
		return nil, fmt.Errorf("Customer %d has no orders", customerID)
	}

	for _, order := range orders {
		if order.CustomerID == customerID {
			customerOrders = append(customerOrders, order)
		}
	}

	if limit > 0 && limit <= len(customerOrders) {
		return customerOrders[:limit], nil
	}

	return customerOrders, nil
}

func (s *OrderService) GetOrderHistory(limit int) ([]models.Order, error) {
	orders := s.storage.GetOrders()
	if len(orders) == 0 {
		return nil, fmt.Errorf("Нет заказов")
	}

	// Сортируем заказы по UpdatedAt (от новых к старым)
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].UpdatedAt.After(orders[j].UpdatedAt)
	})

	if limit > 0 && limit < len(orders) {
		orders = orders[:limit]
	}

	return orders, nil
}

func (s *OrderService) GetReturnedOrders(page, pageSize int) ([]models.Order, error) {
	if page < 1 {
		return nil, fmt.Errorf("номер страницы должен быть больше 0")
	}
	if pageSize < 1 {
		return nil, fmt.Errorf("размер страницы должен быть больше 0")
	}

	orders := s.storage.GetOrders()
	returnedOrders := []models.Order{}

	for _, order := range orders {
		if order.Status == models.StatusReturned {
			returnedOrders = append(returnedOrders, order)
		}
	}

	// индексы для пагинации
	startIndex := (page - 1) * pageSize
	if startIndex >= len(returnedOrders) {
		return []models.Order{}, nil
	}

	endIndex := startIndex + pageSize
	if endIndex > len(returnedOrders) {
		endIndex = len(returnedOrders)
	}

	return returnedOrders[startIndex:endIndex], nil
}

// либо все заказы возможно принять и они добавляются либо ни один заказ не добавляется
func (s OrderService) ImportOrders(path string) error {
	var orders []models.Order
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err := json.Unmarshal(data, &orders); err != nil {
		return fmt.Errorf("failed to unmarshal orders: %w", err)
	}

	for _, order := range orders {
		if _, err := s.storage.FindOrder(order.ID); err == nil {
			return fmt.Errorf("заказ %vуже принят, импорт отклонен", fmt.Sprint(order))
		}

		if time.Now().After(order.StorageUntil) {
			return fmt.Errorf("заказ %vнеккоректен, импорт отклонен", fmt.Sprint(order))
		}
	}

	for _, order := range orders {
		s.storage.AddOrder(order)
	}

	return nil
}
