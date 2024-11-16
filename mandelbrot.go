package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	width, height := 80, 24
	// Define the complex plane boundaries for the fractal
	xmin, ymin, xmax, ymax := -2.0, -1.2, 1.0, 1.2

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			c := complex(x, y)
			// Determine the color based on the Mandelbrot iteration
			fmt.Print(mandelbrot(c))
		}
		fmt.Println()
	}
}

func mandelbrot(c complex128) string {
	const iterations = 1000
	const contrast = 15

	z := c
	for n := 0; n < iterations; n++ {
		if cmplx.Abs(z) > 2 {
			return string('A' + rune(n/contrast)) // creates a gradient with ASCII letters
		}
		z = z*z + c
	}
	return " "
}
