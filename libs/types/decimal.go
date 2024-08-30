package types

import (
	"fmt"
	"log"
	"math"
)

// Decimal (UNTESTED) (SLOW) is a decimal implementation for monetary purposes primarily
type Decimal struct {
	value   uint64
	decimal uint64
}

// NewDecimal creates a new Decimal
func NewDecimal(value uint64, decimal uint64) Decimal {
	return Decimal{value: value, decimal: decimal}
}

// IntToDecimal converts an int into a decimal
func IntToDecimal(valueInt int) Decimal {
	return Decimal{
		value:   uint64(valueInt),
		decimal: 0,
	}
}

// IntToDecimal converts an int into a decimal
func FloatToDecimal(valueFloat float64) Decimal {
	integerPart := math.Floor(valueFloat)
	fractionalPart := valueFloat - integerPart

	var decimalPlaces int
	for fractionalPart != 0 && decimalPlaces < 18 { // Limiting to 18 decimal places
		fractionalPart *= 10
		decimalPlaces++
	}

	combinedValue := uint64(integerPart*math.Pow(10, float64(decimalPlaces)) + fractionalPart)

	return Decimal{
		value:   combinedValue,
		decimal: uint64(decimalPlaces),
	}
}

// Round rounds the decimal to the precision specified using bankers rounding
func (d Decimal) Round(precision uint64) Decimal {
	if precision > d.decimal {
		factor := uint64(math.Pow(10, float64(precision-d.decimal)))
		return NewDecimal(d.value*factor, precision)
	}

	roundingFactor := uint64(math.Pow(10, float64(d.decimal-precision)))
	halfFactor := roundingFactor / 2
	roundedValue := (d.value + halfFactor) / roundingFactor * roundingFactor

	return NewDecimal(roundedValue, precision)
}

func (d Decimal) ToString(precision uint64) string {
	if precision > 100 {
		log.Printf("Warning: using excessively large precision: %v", precision) // TODO: to be replaced by logger injection
	}

	return fmt.Sprintf("%d.%d", d.value, d.Round(precision).decimal)
}
