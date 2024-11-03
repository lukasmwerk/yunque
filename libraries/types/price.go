package types

// Price (UNTESTED) represents a monetary price
type Price struct {
	Value    Decimal
	Currency CurrencyType // For redundancy in case of errors
}

// Rounds a price to its final transaction amount (bankers rounding)
func (p *Price) Round() *Price {
	return &Price{
		Value:    p.Value.Round(decimalPlaces[p.Currency]),
		Currency: p.Currency,
	}
}

func NewPrice(value Decimal, currency CurrencyType) *Price {
	return &Price{
		Value:    value,
		Currency: currency,
	}
}

// ToString shows the string representation (for display purposes only!)
func (p *Price) ToString() string {
	return p.Value.ToString(decimalPlaces[p.Currency])
}

// ResolveDecimals attempts to fix issues with an improperly set decimals value
func (p *Price) ResolveDecimals() {
	// if p.Currency != nil &&
}
