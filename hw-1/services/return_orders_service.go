package services

import (
	"fmt"
	"hw-1/models"
	"time"
)

// принять возвраты клиента
func (s *OrderService) AcceptReturns(customerID uint, orderIDs ...uint) error {
	for _, id := range orderIDs {
		order, err := s.storage.FindOrder(id)
		if err != nil {
			return err
		}

		if order.CustomerID != customerID {
			return fmt.Errorf("%w : %d", ErrOrderNotBelongToCustomer, order.ID)
		}

		if order.Status != models.StatusDelivered {
			return ErrOrderCannotBeReturned
		}

		if time.Now().Sub(order.DeliveredAt) > time.Hour*24*2 {
			return ErrOrderCannotBeReturned
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
