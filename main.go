package main

import (
	"fmt"
	"os"
)

const (
	aspectRatio = 16.0 / 9.0
	imageWidth  = 1024.0
	imageHeight = imageWidth / aspectRatio

	viewportHeight = 2.0
	viewportWidth  = aspectRatio * viewportHeight
	focalLength    = 1.0
)

var (
	origin     = NewVector(0, 0, 0)
	horizontal = NewVector(viewportWidth, 0, 0)
	vertical   = NewVector(0, viewportHeight, 0)
)

func main() {
	lowerLeftCorner := origin.
		Minus(horizontal.Divide(2)).
		Minus(vertical.Divide(2)).
		Minus(NewVector(0, 0, focalLength))

	fmt.Printf("P3\n")
	fmt.Printf("%d %d\n", int(imageWidth), int(imageHeight))
	fmt.Printf("255\n")

	for j := imageHeight - 1; j >= 0; j-- {
		// Progress tracker.
		fmt.Fprintf(os.Stderr, "\rLines remaining: %d", int(j))

		for i := 0; i < imageWidth; i++ {
			x := float64(i) / (imageWidth - 1)
			y := float64(j) / (imageHeight - 1)

			rayDirection := lowerLeftCorner.
				Plus(horizontal.Multiply(x)).
				Plus(vertical.Multiply(y))

			ray := NewRay(origin, rayDirection)
			fmt.Println(determineRayColor(ray).RGB())
		}
	}

	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

// determineRayColor determines the color of the given ray.
func determineRayColor(ray Ray) Color {
	// If sphere is hit, render red.
	if isSphereHit(NewVector(0, 0, -1), 0.5, ray) {
		return NewColor(1, 0, 0)
	}

	// Render the background.
	dir := ray.Direction.Direction()
	t := 0.5 * (dir.Y + 1)
	return NewColor(1, 1, 1).Lerp(NewColor(0.5, 0.7, 1.0), t)
}

// isSphereHit returns true if the given sphere is hit by the given ray.
func isSphereHit(center Vec3, radius float64, ray Ray) bool {
	origin2Center := ray.Origin.Minus(center)

	a := ray.Direction.Dot(ray.Direction)
	bHalf := origin2Center.Dot(ray.Direction)
	c := origin2Center.Dot(origin2Center) - radius*radius

	discriminant := bHalf*bHalf - a*c
	return discriminant >= 0
}
