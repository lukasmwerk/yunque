package models

import "github.com/lukasmwerk/yunque/libraries/types"

// Make represents a manufacturer of a product
type Make struct {
	Id   types.ID
	Name string
}
