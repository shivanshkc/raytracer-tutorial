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

	maxDiffusionDepth = 10
	samplesPerPixel   = 50
)

// cameraOptions holds all the camera configs.
var cameraOptions = CameraOptions{
	LookFrom:            NewVector(13, 2, 3),
	LookAt:              NewVector(0, 0, 0),
	Up:                  NewVector(0, 1, 0),
	AspectRatio:         aspectRatio,
	FieldOfViewVertical: 20,
	Aperture:            0.1,
	FocusDistance:       10,
}

// List of objects that will be rendered.
var hittableGroup = NewHittableGroup([]Hittable{
	Sphere{
		Center: NewVector(0, 1, 0),
		Radius: 1.0,
		Mat: Dielectric{
			RefractiveIndex: 1.5,
		},
	},
	Sphere{
		Center: NewVector(-4, 1, 0),
		Radius: 1,
		Mat: Lambertian{
			Attenuation: NewColor(0.4, 0.2, 0.1),
		},
	},
	Sphere{
		Center: NewVector(4, 1, 0),
		Radius: 1,
		Mat: Metal{
			Attenuation: NewColor(0.7, 0.6, 0.5),
			Fuzz:        0,
		},
	},
	Sphere{
		Center: NewVector(0, -1000, -1),
		Radius: 1000,
		Mat: Lambertian{
			Attenuation: NewColor(0.5, 0.5, 0.5),
		},
	},
})

func main() {
	fmt.Printf("P3\n")
	fmt.Printf("%d %d\n", int(imageWidth), int(imageHeight))
	fmt.Printf("255\n")

	camera := NewCamera(cameraOptions)

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

				ray := camera.CastRay(x, y)

				rayCol := determineRayColor(ray, hittableGroup, maxDiffusionDepth)
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
