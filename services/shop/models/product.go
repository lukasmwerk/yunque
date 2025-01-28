package models

import (
	"github.com/lukasmwerk/yunque/libraries/images"
	"github.com/lukasmwerk/yunque/libraries/types"
)

type Product struct {
	Id          types.ID
	Name        string
	Price       types.Price
	Description string
	Image       images.Image

	Tags ProductTags
}

type ProductTags struct {
	Color        ProductColor
	Manufacturer ProductManufacturer
	Quality      ProductQuality
	Size         ProductSize
	Type         ProductType
}

// ProductDetails contains variant-specific data. For example, a product of a certain
// color may have its own price or description, but we still want to display the parent
// listing that contains the same options lists.
type ProductDetails struct {
	ProductID types.ID
	Price     types.Price
}

type ProductColor struct {
	ID             types.ID // local value, since a color is not part of a predefined set
	Name           string
	Thumbnail      types.Thumbnail
	ProductDetails ProductDetails
}

type ProductManufacturer struct {
	ManufacturerID types.ID
	Name           string
	ProductDetails ProductDetails
}

type ProductQuality struct {
	ID             types.ID // local value
	Name           string
	ProductDetails ProductDetails
}

type ProductSize struct {
	SizeID         types.ID
	Name           string // TODO: add size model or type
	ProductDetails ProductDetails
}

type ProductType struct {
	ID             types.ID
	Name           string
	ProductDetails ProductDetails
}
