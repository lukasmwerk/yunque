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

// Round rounds the decimal to the precision specified
func (d *Decimal) Round(precision uint64) {
	if precision >= d.decimal {
		// No rounding needed if the requested precision is greater than or equal to the current precision
		return
	}

	// Calculate the factor to reduce the precision
	factor := uint64(1)
	for i := uint64(0); i < d.decimal-precision; i++ {
		factor *= 10
	}

	// Calculate the new value and remainder
	scaledValue := d.value / factor
	remainder := d.value % factor

	// Determine if we should round up
	if remainder >= factor/2 {
		scaledValue += 1
	}

	// Update the value and precision
	d.value = scaledValue
	d.decimal = precision
}

func bankersRoundStr(value uint64, precision uint64) string {
	valueStr := fmt.Sprintf("%d", value)
	if len(valueStr) <= int(precision) {
		return fmt.Sprintf("%s%0*d", valueStr, int(precision)-len(valueStr), 0)
	} else {
		return fmt.Sprintf("%s", valueStr[:precision])
	}
}

func (d *Decimal) ToString(precision uint64) string {
	if precision > 100 {
		log.Printf("Warning: using excessively large precision: %v", precision) // TODO: to be replaced by logger injection
	}

	return fmt.Sprintf("%d.%s", d.value, bankersRoundStr(d.decimal, precision))
}
