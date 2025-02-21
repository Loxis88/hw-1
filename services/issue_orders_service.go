package services

import (
	"fmt"
	"hw-1/models"
	"time"
)

// Выдать заказы
func (s *OrderService) IssueOrders(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return fmt.Errorf("cannot deliver orders: %w", err)
		}

		if err := s.isOrderBelongsToCustomer(order, customerID); err != nil {
			return fmt.Errorf("cannot deliver orders: %w", err)
		}

		if order.Status != models.StatusNew {
			return fmt.Errorf("cannot deliver order %d: %w", id, ErrOrderCannotBeDelivered)
		}

		if order.Status == models.StatusNew {
			if time.Now().After(order.StorageUntil) {
				order.Status = models.StatusExpired
				s.storage.UpdateOrder(*order)
				return fmt.Errorf("cannot deliver order %d: %w", id, ErrOrderCannotBeDelivered)
			}

			order.DeliveredAt = time.Now()
			if err := s.updateOrderStatus(order, models.StatusDelivered); err != nil {
				return err
			}
		}

	}

	return nil
}
