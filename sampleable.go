package selection

import (
	"image"
	"image/color"
)

// Sampleable is a selection that can be sampled by color
// and analyzed about its sample
type Sampleable interface {
	GetAverageColorRGBA(img image.Image) *color.RGBA
}

// IsWithinColorRGBA checks whether the average sampled
// color of a Sampleable is within some tolerance of
// the expected color
func IsWithinColorRGBA(
	sampleable Sampleable,
	img image.Image,
	expected *color.RGBA,
	tolerance byte,
) bool {
	actual := sampleable.GetAverageColorRGBA(img)

	return BytesWithinTolerance(actual.R, expected.R, tolerance) && BytesWithinTolerance(actual.G, expected.G, tolerance) && BytesWithinTolerance(actual.B, expected.B, tolerance) && BytesWithinTolerance(actual.A, expected.A, tolerance)
}
