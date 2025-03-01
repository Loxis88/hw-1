package packaging

type FilmPackage struct{}

func (f FilmPackage) Validate(weight float64) error {
	return nil // нет ограничений
}

func (f FilmPackage) CalculateCost(cost float64) float64 {
	return cost + 1
}
