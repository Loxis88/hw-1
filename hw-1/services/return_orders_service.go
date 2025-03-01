package services

import (
	"fmt"
	"hw-1/models"
	"time"
)

// AcceptReturns принимает возвраты клиента
func (s *OrderService) AcceptReturns(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.checkReturnEligibility(id, customerID)
		if err != nil {
			return err
		}

		if err := s.processReturn(order); err != nil {
			return err
		}
	}
	return nil
}

// checkReturnEligibility проверяет, можно ли вернуть заказ
func (s *OrderService) checkReturnEligibility(orderID, customerID uint) (*models.Order, error) {
	order, err := s.storage.FindOrder(orderID)
	if err != nil {
		return nil, err
	}

	if order.CustomerID != customerID {
		return nil, fmt.Errorf("%w : %d", ErrOrderNotBelongToCustomer, order.ID)
	}

	if order.Status != models.StatusDelivered || time.Now().Sub(order.DeliveredAt) > 48*time.Hour {
		return nil, ErrOrderCannotBeReturned
	}

	return order, nil
}

// processReturn обновляет статус заказа на возвращённый
func (s *OrderService) processReturn(order *models.Order) error {
	return s.updateOrderStatus(order, models.StatusReturned)
}
