package selection

import "image/draw"

// PixelArray is a selection of pixels
type PixelArray []Pixel

// Touch touches a random pixel in the array
func (pa PixelArray) Touch(medium Medium, rand Random) error {
	return pa[rand.Intn(len(pa))].Touch(medium, rand)
}

// Draw draws the pixels on an image
func (pa PixelArray) Draw(img draw.Image) {
	for _, pixel := range pa {
		pixel.Draw(img)
	}
}
