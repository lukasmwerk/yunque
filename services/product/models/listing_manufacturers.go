package models

import "github.com/lukasmwerk/yunque/libs/types"

// ListingManufacturers implements brand options for browsing alternative manufacturers
type ListingManufacturers struct {
	List []*ListingManufacturer
}

type ListingManufacturer struct {
	Name types.ManufacturerName
}
