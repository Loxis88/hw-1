package services

import (
	"encoding/json"
	"fmt"
	"hw-1/models"
	"os"
	"time"
)

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
