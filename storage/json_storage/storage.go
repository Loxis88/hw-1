package json_storage

import (
	"encoding/json"
	"fmt"
	"os"

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
		return nil, err
	}

	var orders []models.Order

	err = json.Unmarshal(file, &orders)
	if err != nil {
		return nil, err
	}

	return &Storage{
		orders: orders,
		path:   path,
	}, nil
}

func (s *Storage) save() error {
	file, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(s.orders)
}

func (s *Storage) AddOrder(order models.Order) error {
	s.orders = append(s.orders, order)
	return s.save()
}

func (s *Storage) DeleteOrder(id uint) error {
	for i, order := range s.orders {
		if order.ID == id {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			s.save()
			return nil
		}
	}
	return fmt.Errorf("order with id %d not found", id)
}

func (s *Storage) GetOrders() ([]models.Order, error) {
	return s.orders, nil
}

func (s *Storage) FindOrder(id uint) (*models.Order, error) {
	for _, order := range s.orders {
		if order.ID == id {
			return &order, nil
		}
	}
	return nil, fmt.Errorf("order with id %d not found", id)
}

func (s *Storage) Test() {
	for _, ordeer := range s.orders {
		fmt.Println(ordeer)
	}
}
