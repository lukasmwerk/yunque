package models

import "github.com/lukasmwerk/yunque/libs/types"

type Listing struct {
	Name    string
	Price   types.PriceRange
	Options ListingOptions
}
