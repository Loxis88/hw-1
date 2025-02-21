package services

import (
	"fmt"
	"hw-1/models"
	"time"
)

func (s *OrderService) ReturnOrderToCourier(orderID uint) error {
	order, err := s.storage.FindOrder(orderID)
	if err != nil {
		return err
	}

	if order.Status == models.StatusExpired {
		return s.storage.DeleteOrder(orderID)
	}

	if order.Status == models.StatusNew {
		// проверка на тот случай если заказ был не просрочен на момент послежнего обновления хранилища, но к моменту выполнения операции просрочился
		if time.Now().After(order.StorageUntil) {
			return s.storage.DeleteOrder(orderID)
		}
		return fmt.Errorf("order %d: %w", orderID, ErrStoragePeriodNotExpired)
	}

	if order.Status == models.StatusDelivered {
		return fmt.Errorf("order %d: %w", orderID, ErrOrderCannotBeReturnToCurier)
	}

	return s.storage.DeleteOrder(orderID)
}
