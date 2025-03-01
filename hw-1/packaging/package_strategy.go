package packaging

import "fmt"

type BagPackage struct{}

func (b BagPackage) Validate(weight float64) error {
	if weight >= 10 {
		return fmt.Errorf("bag package cannot handle weight >= 10 kg")
	}
	return nil
}

func (b BagPackage) CalculateCost(cost float64) float64 {
	return cost + 5
}
