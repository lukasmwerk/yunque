package images

import "github.com/lukasmwerk/yunque/libraries/types"

type Image struct {
	Id           types.ID
	Name         string
	Description  string
	FullURL      string // The original unscaled image
	ScaledURL    string // A smaller compressed/scaled image for display
	ThumbnailUrl string // A thumbnail for the image
}
