package packaging

import (
	"fmt"
	"hw-1/models"
)

type PackageStrategy interface {
	Validate(weight float64) error
	CalculateCost(cost float64) float64
}

type PackageContext struct {
	strategies map[models.PackageType]PackageStrategy
}

func NewPackageContext() *PackageContext {
	return &PackageContext{
		strategies: map[models.PackageType]PackageStrategy{
			models.PackageTypePackage: BagPackage{},
			models.PackageTypeBox:     BoxPackage{},
			models.PackageTypeFilm:    FilmPackage{},
		},
	}
}

func (c *PackageContext) GetStrategy(packageType models.PackageType) (PackageStrategy, error) {
	strategy, ok := c.strategies[packageType]
	if !ok {
		return nil, fmt.Errorf("unknown package type: %s", packageType)
	}
	return strategy, nil
}
