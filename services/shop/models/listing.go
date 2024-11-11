package models

import "github.com/lukasmwerk/yunque/libraries/types"

type Listing struct {
	Name    string
	Price   types.PriceRange
	Options ListingOptions
}

// colorNameString is a string wrapper to enforce proper naming
type ColorNameString string

// ToString (NOT DONE)
func (s ColorNameString) ToString() string {
	return string(s)
}

// NewColorNameString (NOT DONE)
func NewColorNameString(s string) (string, error) {
	return "", nil
}

// ListingColor implements color options for products
type ListingColors struct {
	List []*ListingColor // First item is default in cases where that is relevant
}

// ListingColor (UNTESTED) is a color of a product listed in an item listing
type ListingColor struct {
	Name      ColorNameString
	Thumbnail types.Thumbnail
	Product   Product
}

// ListingManufacturers implements brand options for browsing alternative manufacturers
type ListingManufacturers struct {
	List []*ListingManufacturer
}

type ListingManufacturer struct {
	Name types.ManufacturerName
}

// ListingOptions (UNTESTED) is a wrapper struct for selectable options such as size, color, etc
type ListingOptions struct {
	Color        ListingColors
	Manufacturer ListingManufacturers
	Quality      ListingQualities
	Size         ListingSizes
	Type         ListingTypes
}

// ListingQualities implements quality options for products such as refurbished
type ListingQualities struct{}

type ListingSizes struct {
}

type ListingTypes struct {
}
