package models

import "github.com/lukasmwerk/yunque/libs/types"

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
