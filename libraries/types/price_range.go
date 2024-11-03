package types

// PriceRange is a wrapper for Prices enablng price range functionality
type PriceRange struct {
	Low  Price // Or only price
	High Price // Empty if only one price
}
