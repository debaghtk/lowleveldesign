package shipping

type Package struct {
	weight   float64
	distance float64
}

type ShippingCalculatorInterface interface {
	CalculateCost(Package) float64
}

type ExpressShippingCalculator struct{}
type RegularShippingCalculator struct{}
type InternationalShippingCalculator struct{}
type OvernightShippingCalculator struct{}

func (rsc *RegularShippingCalculator) CalculateCost(pkg Package) float64 {
	return pkg.weight * 1.5 * pkg.distance // Regular rate
}

func (esc *ExpressShippingCalculator) CalculateCost(pkg Package) float64 {
	return pkg.weight * 2.5 * pkg.distance * 1.5 // Express rate
}

func (isc *InternationalShippingCalculator) CalculateCost(pkg Package) float64 {
	return pkg.weight * pkg.distance * 5
}

func (osc *OvernightShippingCalculator) CalculateCost(pkg Package) float64 {
	return pkg.weight * pkg.distance * 4
}
