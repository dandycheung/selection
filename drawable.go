package selection

import (
	"image/color"
	"image/draw"
)

var (
	// DrawableColor is the main color of selection drawing
	DrawableColor = color.RGBA{0, 255, 0, 255}
)

// Drawable is a selection that can draw on an image
type Drawable interface {
	Draw(img draw.Image)
}

// DrawBatch draws an array of Drawables on an image
func DrawBatch(drawables []Drawable, img draw.Image) {
	for _, drawable := range drawables {
		drawable.Draw(img)
	}
}
