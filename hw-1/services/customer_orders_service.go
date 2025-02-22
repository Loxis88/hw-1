package services

import (
	"fmt"
	"hw-1/models"
)

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
