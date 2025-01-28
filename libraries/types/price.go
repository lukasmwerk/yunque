package types

// Price represents a monetary price
type Price struct {
	// Value is the numeric representation of the price.
	// Always in USD multiplied by 100, so 2.99 is 299.
	// Note: interested to see what issues I run into
	// if I need to switch to a larger multiplier later.
	Value uint64
}

//// Round rounds a price to its final transaction amount (uses bankers rounding).
//func (p *Price) Round() *Price {
//	return &Price{
//		Value: p.Value.Round(decimalPlaces[p.Currency]),
//	}
//}

// NewPrice creates a new price.
func NewPrice(value uint64, currency CurrencyType) *Price {
	return &Price{
		Value: value,
	}
}

// NewPriceFromFloat creates a price from a float price, for example 3.31
// Extra decimal places beyond the hundreds are shaved off (rounded down)
func NewPriceFromFloat(value float64, currency CurrencyType) *Price {
	return &Price{
		Value: uint64(value * 100),
	}
}

// ToFloat returns a float representation of a price, e.g. 2.99
// Calculations on prices should not be done using floats, this
// function is primarily for display purposes.
func (p *Price) ToFloat() float64 {
	return float64(p.Value / 100)
}
