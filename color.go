package main

import (
	"fmt"
)

// Color represents an RGB color.
type Color struct {
	R float64
	G float64
	B float64
}

// NewColor returns a new color instance.
func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

// RGB converts the color into an "<r> <g> <b>" formatted string that can be used in a PPM file.
func (c Color) RGB() string {
	return fmt.Sprintf("%d %d %d", int(255.999*c.R), int(255.999*c.G), int(255.999*c.B))
}
