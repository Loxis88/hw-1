package services

import (
	"fmt"
	"hw-1/models"
	"sort"
)

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
