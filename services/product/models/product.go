package models

import "github.com/lukasmwerk/yunque/libraries/types"

// Product represents a single product at the lowest layer of abstraction
type Product struct {
	Name  string
	Price types.Price
}
