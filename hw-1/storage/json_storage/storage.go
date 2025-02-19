package json_storage

import (
	"encoding/json"
	"fmt"
	"os"
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

	storage := &Storage{
		orders: orders,
		path:   path,
	}

	storage.ValidateOrders()
	return storage, nil
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

func (s *Storage) ValidateOrders() {
	now := time.Now()
	for i, order := range s.orders {
		if order.Status == models.StatusNew && now.After(order.StorageUntil) {
			s.orders[i].Status = models.StatusExpired
			s.orders[i].UpdatedAt = now
		}
	}
	s.save()
}

// AddOrder adds a new order to the storage
func (s *Storage) AddOrder(order models.Order) error {
	order.UpdatedAt = time.Now()
	if order.Status == "" {
		order.Status = models.StatusNew
	}
	s.orders = append(s.orders, order)
	if err := s.save(); err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}
	return nil
}

// UpdateOrder updates an existing order in the storage
func (s *Storage) UpdateOrder(order models.Order) error {
	for i, o := range s.orders {
		if o.ID == order.ID {
			s.orders[i] = order
			if err := s.save(); err != nil {
				return fmt.Errorf("failed to save order: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("order with id %d not found", order.ID)
}

// DeleteOrder deletes an order from the storage
func (s *Storage) DeleteOrder(id uint) error {
	for i, order := range s.orders {
		if order.ID == id {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			if err := s.save(); err != nil {
				return fmt.Errorf("failed to save order: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("order with id %d not found", id)
}

// GetOrders retrieves all orders from the storage
func (s *Storage) GetOrders() []models.Order {
	return s.orders
}

// FindOrder finds an order by its ID
func (s *Storage) FindOrder(id uint) (*models.Order, error) {
	for i, order := range s.orders {
		if order.ID == id {
			return &s.orders[i], nil
		}
	}
	return nil, fmt.Errorf("order with id %d not found", id)
}
