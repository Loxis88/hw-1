package services

import (
	"fmt"
	"hw-1/models"
	"time"
)

// IssueOrders delivers orders for a customer
func (s *OrderService) IssueOrders(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return fmt.Errorf("cannot deliver orders: %w", err)
		}

		if err := s.checkOrderEligibility(order, customerID); err != nil {
			return fmt.Errorf("cannot deliver order %d: %w", id, err)
		}

		if err := s.processOrder(order); err != nil {
			return fmt.Errorf("cannot deliver order %d: %w", id, err)
		}
	}
	return nil
}

// checkOrderEligibility verifies if order can be delivered
func (s *OrderService) checkOrderEligibility(order *models.Order, customerID uint) error {
	if err := s.isOrderBelongsToCustomer(order, customerID); err != nil {
		return err
	}
	if order.Status != models.StatusNew {
		return ErrOrderCannotBeDelivered
	}
	return nil
}

// processOrder handles the delivery process
func (s *OrderService) processOrder(order *models.Order) error {
	if time.Now().After(order.StorageUntil) {
		order.Status = models.StatusExpired
		s.storage.UpdateOrder(*order)
		return ErrOrderCannotBeDelivered
	}

	order.DeliveredAt = time.Now()
	return s.updateOrderStatus(order, models.StatusDelivered)
}
