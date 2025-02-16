package json_storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"hw-1/models"
)

type Storage struct {
	orders []models.Order
	path   string
}

func New(path string) (*Storage, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Storage{
				orders: []models.Order{},
				path:   path,
			}, nil
		}
		return nil, fmt.Errorf("failed to read storage file: %w", err)
	}

	var orders []models.Order
	if err := json.Unmarshal(file, &orders); err != nil {
		return nil, fmt.Errorf("failed to unmarshal orders: %w", err)
	}

	return &Storage{
		orders: orders,
		path:   path,
	}, nil
}

func (s *Storage) save() error {
	file, err := os.Create(s.path)
	if err != nil {
		return fmt.Errorf("failed to create storage file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(s.orders); err != nil {
		return fmt.Errorf("failed to encode orders: %w", err)
	}
	return nil
}

// AddOrder добавляет новый заказ
func (s *Storage) AddOrder(order models.Order) error {
	if _, err := s.FindOrder(order.ID); err == nil {
		return fmt.Errorf("order with id %d already exists", order.ID)
	}

	s.orders = append(s.orders, order)
	return s.save()
}

func (s *Storage) UpdateOrder(order models.Order) error {
	for i, o := range s.orders {
		if o.ID == order.ID {
			s.orders[i] = order
			return s.save()
		}
	}
	return fmt.Errorf("order with id %d not found", order.ID)
}

func (s *Storage) DeleteOrder(id uint) error {
	for i, order := range s.orders {
		if order.ID == id {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("order with id %d not found", id)
}

func (s *Storage) GetOrders() []models.Order {
	return s.orders
}

func (s *Storage) GetOrdersByCustomer(customerID uint, lastN int, inStorageOnly bool) []models.Order {
	var result []models.Order

	for _, order := range s.orders {
		if order.CustomerID == customerID {
			if inStorageOnly && order.Status != models.StatusNew {
				continue
			}
			result = append(result, order)
		}
	}

	if lastN > 0 && len(result) > lastN {
		return result[len(result)-lastN:]
	}

	return result
}

func (s *Storage) FindOrder(id uint) (*models.Order, error) {
	for i, order := range s.orders {
		if order.ID == id {
			return &s.orders[i], nil
		}
	}
	return nil, fmt.Errorf("order with id %d not found", id)
}

func (s *Storage) GetExpiredOrders() []models.Order {
	var expired []models.Order
	now := time.Now()

	for _, order := range s.orders {
		if order.Status == models.StatusNew && now.After(order.StorageUntil) {
			expired = append(expired, order)
		}
	}

	return expired
}

func (s *Storage) GetReturnedOrders() []models.Order {
	orders := s.GetOrders()

	var result []models.Order
	for _, order := range orders {
		if order.Status == models.StatusReturned {
			result = append(result, order)
		}
	}
	return result
}

func (s *Storage) GetOrdersHistory(limit int) ([]models.Order, error) {
	// Получаем все заказы
	orders := s.GetOrders()

	// Сортируем по времени обновления (в порядке убывания)
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].UpdatedAt.After(orders[j].UpdatedAt)
	})

	// Применяем лимит
	if len(orders) > limit {
		orders = orders[:limit]
	}

	return orders, nil
}
