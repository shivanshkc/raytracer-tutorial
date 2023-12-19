package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	aspectRatio = 16.0 / 9.0
	imageWidth  = 800.0
	imageHeight = imageWidth / aspectRatio

	viewportHeight = 2.0
	viewportWidth  = aspectRatio * viewportHeight
	focalLength    = 1.0

	maxDiffusionDepth = 10
	// For anti-aliasing.
	samplesPerPixel = 100
)

var (
	origin     = NewVector(0, 0, 0)
	horizontal = NewVector(viewportWidth, 0, 0)
	vertical   = NewVector(0, viewportHeight, 0)

	// List of objects that will be rendered.
	hittableGroup = NewHittableGroup([]Hittable{
		Sphere{
			Center: NewVector(0, 0, -1),
			Radius: 0.5,
			Mat: Lambertian{
				Attenuation: NewColor(0.7, 0.3, 0.3),
			},
		},
		Sphere{
			Center: NewVector(-1, 0, -1),
			Radius: 0.5,
			Mat: Metal{
				Attenuation: NewColor(0.8, 0.8, 0.8),
				Fuzz:        0.3,
			},
		},
		Sphere{
			Center: NewVector(1, 0, -1),
			Radius: 0.5,
			Mat: Metal{
				Attenuation: NewColor(0.8, 0.6, 0.2),
				Fuzz:        0,
			},
		},
		Sphere{
			Center: NewVector(0, -100.5, -1),
			Radius: 100,
			Mat: Lambertian{
				Attenuation: NewColor(0.8, 0.8, 0),
			},
		},
	})
)

func main() {
	lowerLeftCorner := origin.
		Minus(horizontal.Divide(2)).
		Minus(vertical.Divide(2)).
		Minus(NewVector(0, 0, focalLength))

	fmt.Printf("P3\n")
	fmt.Printf("%d %d\n", int(imageWidth), int(imageHeight))
	fmt.Printf("255\n")

	// Create a random number generator.
	var randSeed int64 = int64(time.Now().Nanosecond())
	randomGen := rand.New(rand.NewSource(randSeed))

	for j := imageHeight - 1; j >= 0; j-- {
		// Progress tracker.
		fmt.Fprintf(os.Stderr, "\rLines remaining: %d", int(j))

		for i := 0.0; i < imageWidth; i++ {
			color := NewColor(0, 0, 0)

			for s := 0; s < samplesPerPixel; s++ {
				x := (i + randomGen.Float64()) / (imageWidth - 1)
				y := (j + randomGen.Float64()) / (imageHeight - 1)

				rayDirection := lowerLeftCorner.
					Plus(horizontal.Multiply(x)).
					Plus(vertical.Multiply(y))

				rayCol := determineRayColor(NewRay(origin, rayDirection), hittableGroup, maxDiffusionDepth)
				color = NewColor(
					color.R+rayCol.R,
					color.G+rayCol.G,
					color.B+rayCol.B,
				)
			}

			fmt.Println(color.RGB(samplesPerPixel))
		}
	}

	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

// determineRayColor determines the color of the given ray.
func determineRayColor(ray Ray, object Hittable, depth int) Color {
	if depth < 1 {
		return NewColor(0, 0, 0)
	}

	if info, isHit := object.IsHit(ray, 0.001, math.MaxFloat64); isHit {
		scattered, attenuation, isScattered := info.Mat.Scatter(ray, info)
		if !isScattered {
			return NewColor(0, 0, 0)
		}

		scatteredRayCol := determineRayColor(scattered, object, depth-1)
		return NewColor(
			scatteredRayCol.R*attenuation.R,
			scatteredRayCol.G*attenuation.G,
			scatteredRayCol.B*attenuation.B,
		)
	}

	// Render the background.
	dir := ray.Direction.Direction()
	t := 0.5 * (dir.Y + 1)
	return NewColor(1, 1, 1).Lerp(NewColor(0.5, 0.7, 1.0), t)
}
