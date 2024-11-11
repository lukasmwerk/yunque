package models

import (
	"github.com/lukasmwerk/yunque/libraries/images"
	"github.com/lukasmwerk/yunque/libraries/types"
)

type Product struct {
	Id          types.ID
	Name        string
	Make        Make
	Description string
	Price       types.Price
	Image       images.Image
}
