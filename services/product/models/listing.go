package models

import "github.com/lukasmwerk/yunque/libraries/types"

type Listing struct {
	Name    string
	Price   types.PriceRange
	Options ListingOptions
}
