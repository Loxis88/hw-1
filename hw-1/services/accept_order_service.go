package services

import (
	"hw-1/models"
	"hw-1/packaging"
	"time"
)

// Принять заказ от курьера
func (s *OrderService) AcceptOrder(orderID uint, customerID uint, storageDate time.Time, weight float64, cost float64, packageType models.PackageType, wrap bool) error {
	order, _ := s.storage.FindOrder(orderID)
	if order != nil {
		return ErrOrderAlreadyExists
	}

	// валидация времени хранения заказа
	if storageDate.Before(time.Now()) {
		return ErrInvalidStorageDate
	}

	if packageType != "" {
		PackageRegistry := packaging.NewPackageContext()

		strategy, err := PackageRegistry.GetStrategy(packageType)
		if err != nil {
			return err
		}

		// Если нужна дополнительная обертка пленкой и это не сама пленка
		if wrap && packageType != models.PackageTypeFilm {
			strategy = packaging.NewFilmDecorator(strategy)
		}

		err = strategy.Validate(weight)
		if err != nil {
			return err
		}
		cost = strategy.CalculateCost(cost)
	}

	newOrder := models.Order{
		ID:           orderID,
		CustomerID:   customerID,
		StorageUntil: storageDate,
		Status:       models.StatusNew,
		UpdatedAt:    time.Now(),
		Weight:       weight,
		Cost:         cost,
		PackageType:  packageType,
	}

	return s.storage.AddOrder(newOrder)
}
