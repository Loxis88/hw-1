package packaging

import (
	"fmt"
	"hw-1/models"
)

// PackageStrategy - общий интерфейс для всех упаковок
type PackageStrategy interface {
	Validate(weight float64) error
	CalculateCost(cost float64) float64
}

// PackageContext - контекст для стратегий, по хотя по сути регистр со всеми стратегиями
type PackageRegistry struct {
	strategies map[models.PackageType]PackageStrategy
}

func NewPackageContext() *PackageRegistry {
	return &PackageRegistry{
		strategies: map[models.PackageType]PackageStrategy{
			models.PackageTypePackage: BagPackage{},
			models.PackageTypeBox:     BoxPackage{},
			models.PackageTypeFilm:    FilmPackage{},
		},
	}
}

func (c *PackageRegistry) GetStrategy(packageType models.PackageType) (PackageStrategy, error) {
	strategy, ok := c.strategies[packageType]
	if !ok {
		return nil, fmt.Errorf("unknown package type: %s", packageType)
	}
	return strategy, nil
}

// PackageDecorator - декоратор для упаковки
type PackageDecorator struct {
	wrapped PackageStrategy
}

// FilmDecorator - конкретный декоратор для пленки
type FilmDecorator struct {
	PackageDecorator
}

// NewFilmDecorator создает новый экземпляр декоратора пленки
func NewFilmDecorator(wrapped PackageStrategy) *FilmDecorator {
	return &FilmDecorator{
		PackageDecorator: PackageDecorator{wrapped: wrapped},
	}
}

// Validate проверяет, может ли упаковка с пленкой выдержать указанный вес
func (f *FilmDecorator) Validate(weight float64) error {
	// Пленка не накладывает дополнительных ограничений
	return f.wrapped.Validate(weight)
}

// CalculateCost вычисляет стоимость общей упаковки с пленкой
func (f *FilmDecorator) CalculateCost(cost float64) float64 {

	baseCost := f.wrapped.CalculateCost(cost)
	return baseCost + 1
}
