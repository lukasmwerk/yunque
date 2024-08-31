package models

import "github.com/lukasmwerk/yunque/libs/types"

// Product represents a single product at the lowest layer of abstraction
type Product struct {
	Name  string
	Price types.Price
}
