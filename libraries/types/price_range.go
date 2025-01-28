package types

// PriceRange is a wrapper for Prices enablng price range functionality
type PriceRange struct {
	Low  Price // Or only price
	High Price // Empty if only one price
}

func NewPriceRange() *PriceRange {
	return &PriceRange{}
}

// Include adds a price to the PriceRange. If the value is not a potential
// low or high, it will be ignored, meaning some failures are quiet.
// A valid price range must be passed.
func (pr *PriceRange) Include(price Price) {
	if price.Value < pr.Low.Value || pr.Low.Value == 0 {
		pr.Low = price
	}
	if price.Value > pr.High.Value {
		pr.High = price
	}
}
