package models

// ProductOptions (UNTESTED) is a wrapper struct for selectable options such as size, color, etc
type ProductOptions struct {
	Color        ProductColor
	Maunfacturer ProductManufacturer
	Quality      ProductQuality
	Size         ProductSize
	Type         ProductType
}
