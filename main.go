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
			r, g, b := float64(i)/(imageWidth-1), float64(j)/(imageHeight-1), 0.25
			fmt.Println(NewColor(r, g, b).RGB())
		}
	}

	fmt.Fprintf(os.Stderr, "\nDone.\n")
}
