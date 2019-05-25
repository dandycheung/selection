package selection

import (
	"fmt"
	"image/draw"
)

// GridTemplate is a configuration template
// to create a new Grid selection
type GridTemplate struct {
	TopLeft       *Pixel
	Rows          int
	Cols          int
	PaddingWidth  int
	PaddingHeight int
	CellWidth     int
	CellHeight    int
}

// Grid is a selection of rectangles in a grid
type Grid struct {
	Rects []*Rect
	rows  int
	cols  int
}

// NewGrid creates a Grid selection out of a
// GridTemplate
func NewGrid(template *GridTemplate) *Grid {
	rects := make([]*Rect, template.Rows*template.Cols)

	for i := range rects {
		x := i % template.Cols
		y := i / template.Cols

		topLeftX := template.TopLeft.X + (x * template.CellWidth) + template.PaddingWidth
		topLeftY := template.TopLeft.Y + (y * template.CellHeight) + template.PaddingHeight

		rects[i] = &Rect{
			TopLeft: &Pixel{topLeftX, topLeftY},
			Width:   template.CellWidth - (2 * template.PaddingWidth),
			Height:  template.CellHeight - (2 * template.PaddingHeight),
		}
	}

	return &Grid{
		Rects: rects,
		rows:  template.Rows,
		cols:  template.Cols,
	}
}

// TouchXY touches a specific Rect in the grid with
// the top left Rect being 0,0
func (g *Grid) TouchXY(medium Medium, rand Random, x int, y int) error {
	index := y*g.cols + x

	if index < 0 || index >= len(g.Rects) {
		return fmt.Errorf("coordinates out of range: %d,%d", x, y)
	}

	return g.Rects[index].Touch(medium, rand)
}

// TouchIndex touches a Rect in the grid at
// a specific index, with the top left Rect being
// index 0 with the index increasing left to right
func (g *Grid) TouchIndex(medium Medium, rand Random, index int) error {
	if index < 0 || index >= len(g.Rects) {
		return fmt.Errorf("index out of range: %d", index)
	}

	return g.Rects[index].Touch(medium, rand)
}

// Draw draws the grid of Rects on an image
func (g *Grid) Draw(img draw.Image) {
	for _, r := range g.Rects {
		r.Draw(img)
	}
}
