package selection

import (
	"image"
	"image/color"
	"image/draw"
)

// Pixel is a selection of one pixel
type Pixel struct {
	X int
	Y int
}

// Touch touches the pixel on the medium
func (p *Pixel) Touch(medium Medium, rand Random) error {
	return medium.Touch(p.X, p.Y)
}

// Draw draws the Pixel on an image
func (p *Pixel) Draw(img draw.Image) {
	img.Set(p.X, p.Y, DrawableColor)
}

// GetAverageColorRGBA gets the average sampled color
// from an image at the Pixel
func (p *Pixel) GetAverageColorRGBA(img image.Image) *color.RGBA {
	red, green, blue, alpha := img.At(p.X, p.Y).RGBA()

	return &color.RGBA{
		byte(red >> 8),
		byte(green >> 8),
		byte(blue >> 8),
		byte(alpha >> 8),
	}
}
