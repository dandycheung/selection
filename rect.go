package selection

import (
	"image"
	"image/color"
	"image/draw"
)

// Rect is a rectangle selection
type Rect struct {
	TopLeft *Pixel
	Width   int
	Height  int
}

// Touch touches a random pixel within the Rect
// on the medium
func (r *Rect) Touch(medium Medium, rand Random) error {
	return medium.Touch(r.TopLeft.X+rand.Intn(r.Width), r.TopLeft.Y+rand.Intn(r.Height))
}

// Draw draws the Rect on an image
func (r *Rect) Draw(img draw.Image) {
	for x := r.TopLeft.X; x < r.TopLeft.X+r.Width; x += r.Width - 1 {
		for y := r.TopLeft.Y; y < r.TopLeft.Y+r.Height; y++ {
			img.Set(x, y, DrawableColor)
		}
	}

	for x := r.TopLeft.X; x < r.TopLeft.X+r.Width; x++ {
		for y := r.TopLeft.Y; y < r.TopLeft.Y+r.Height; y += r.Height - 1 {
			img.Set(x, y, DrawableColor)
		}
	}
}

// GetAverageColorRGBA gets the average sampled color
// from an image within the Rect
func (r *Rect) GetAverageColorRGBA(img image.Image) *color.RGBA {
	var reds int
	var greens int
	var blues int
	var alphas int
	numPixels := r.Width * r.Height

	for x := r.TopLeft.X; x < r.TopLeft.X+r.Width; x++ {
		for y := r.TopLeft.Y; y < r.TopLeft.Y+r.Height; y++ {
			red, green, blue, alpha := img.At(r.TopLeft.X, r.TopLeft.Y).RGBA()

			reds += int(red >> 8)
			greens += int(green >> 8)
			blues += int(blue >> 8)
			alphas += int(alpha >> 8)
		}
	}

	return &color.RGBA{
		byte(reds / numPixels),
		byte(greens / numPixels),
		byte(blues / numPixels),
		byte(alphas / numPixels),
	}
}
