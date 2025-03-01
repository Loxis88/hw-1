package services

import (
	"fmt"
	"time"

	"hw-1/models"
	"hw-1/storage"
)

type OrderServiceInterface interface {
	AcceptOrder(orderID uint, customerID uint, storageDate time.Time, weight float64, cost float64, packageType models.PackageType, wrap bool) error
	ReturnOrderToCourier(orderID uint) error
	IssueOrders(customerID uint, orderIDs ...uint) error
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

// функции бизнес логики были разнесены в отдельные фаилы в самом пакете services
// в service.go остались небольшие вспомогательные функции

func (s *OrderService) isOrderBelongsToCustomer(order *models.Order, customerID uint) error {
	if order.CustomerID != customerID {
		return fmt.Errorf("%w : %d", ErrOrderNotBelongToCustomer, order.ID)
	}
	return nil
}

func (s *OrderService) updateOrderStatus(order *models.Order, status models.OrderStatus) error {
	order.Status = status
	order.UpdatedAt = time.Now()
	return s.storage.UpdateOrder(*order)
}
