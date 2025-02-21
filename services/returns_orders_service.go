package services

import (
	"fmt"
	"hw-1/models"
)

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
