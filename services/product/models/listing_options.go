package models

// ListingOptions (UNTESTED) is a wrapper struct for selectable options such as size, color, etc
type ListingOptions struct {
	Color        ListingColors
	Maunfacturer ListingManufacturers
	Quality      ListingQualities
	Size         ListingSizes
	Type         ListingTypes
}
