package main

import (
	"fmt"
	"os"
)

const (
	imageWidth  = 256
	imageHeight = 256
)

func main() {
	fmt.Printf("P3\n")
	fmt.Printf("%d %d\n", imageWidth, imageHeight)
	fmt.Printf("255\n")

	for j := imageHeight - 1; j >= 0; j-- {
		// Progress tracker.
		fmt.Fprintf(os.Stderr, "\rLines remaining: %d", j)

		for i := 0; i < imageWidth; i++ {
			r, g, b := float32(i)/(imageWidth-1),
				float32(j)/(imageHeight-1), 0.25

			rInt, gInt, bInt := int(255.999*r),
				int(255.999*g), int(255.999*b)

			fmt.Printf("%d %d %d\n", rInt, gInt, bInt)
		}
	}

	fmt.Fprintf(os.Stderr, "\nDone.\n")
}
