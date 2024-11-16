package main

import (
	"fmt"
)

func main() {
	height := 32 // Height of the triangle, can be adjusted to powers of 2 (e.g., 16, 32, 64)

	// Initialize a blank grid to hold the triangle
	triangle := make([][]rune, height)
	for i := range triangle {
		triangle[i] = make([]rune, 2*height-1)
		for j := range triangle[i] {
			triangle[i][j] = ' ' // Fill with spaces
		}
	}

	// Draw the Sierpinski triangle
	drawSierpinski(triangle, height-1, height-1, height)

	// Print the triangle
	for _, row := range triangle {
		fmt.Println(string(row))
	}
}

// Recursive function to draw the Sierpinski triangle
func drawSierpinski(triangle [][]rune, row, col, size int) {
	if size < 1 || row < 0 || col < 0 || row >= len(triangle) || col >= len(triangle[row]) {
		return // Prevent out-of-bounds access
	}
	if size == 1 {
		triangle[row][col] = '▲' // Draw the top of a triangle
		return
	}

	// Recursive calls for the three corners of the current triangle
	half := size / 2
	drawSierpinski(triangle, row, col, half)           // Top triangle
	drawSierpinski(triangle, row+half, col-half, half) // Bottom-left triangle
	drawSierpinski(triangle, row+half, col+half, half) // Bottom-right triangle
}
