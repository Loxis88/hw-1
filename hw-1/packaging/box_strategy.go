package packaging

import "fmt"

type BoxPackage struct{}

func (b BoxPackage) Validate(weight float64) error {
    if weight >= 30 {
        return fmt.Errorf("box package cannot handle weight >= 30 kg")
    }
    return nil
}

func (b BoxPackage) CalculateCost(cost float64) float64 {
    return cost + 20
}
