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

// NewColorFromVec3 converts the given Vec3 into a Color.
func NewColorFromVec3(vec Vec3) Color {
	return Color{R: vec.X, G: vec.Y, B: vec.Z}
}

// RGB converts the color into an "<r> <g> <b>" formatted string that can be used in a PPM file.
func (c Color) RGB(samplesPerPixel int) string {
	scale := 1.0 / float64(samplesPerPixel)

	return fmt.Sprintf(
		"%d %d %d",
		int(256*clamp(c.R*scale, 0, 0.9999)),
		int(256*clamp(c.G*scale, 0, 0.9999)),
		int(256*clamp(c.B*scale, 0, 0.9999)),
	)
}

// Lerp stand for linear interpolation.
func (c Color) Lerp(end Color, factor float64) Color {
	oneMinusFactor := 1 - factor

	return Color{
		oneMinusFactor*c.R + factor*end.R,
		oneMinusFactor*c.G + factor*end.G,
		oneMinusFactor*c.B + factor*end.B,
	}
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
