package main

import (
	"fmt"
	"math"
)

// drawTree recursively draws a tree pattern
func drawTree(x, y, length int, angle float64, depth int) {
	if depth == 0 || length < 1 {
		return
	}

	// Calculate the end point of the branch
	endX := x + int(float64(length)*math.Cos(angle))
	endY := y - int(float64(length)*math.Sin(angle))

	// Draw the current branch
	drawLine(x, y, endX, endY)

	// Recursive calls for left and right branches
	drawTree(endX, endY, length/2, angle-math.Pi/4, depth-1)
	drawTree(endX, endY, length/2, angle+math.Pi/4, depth-1)
}

// drawLine draws a line between two points in the terminal
func drawLine(x1, y1, x2, y2 int) {
	dx := int(math.Abs(float64(x2 - x1)))
	dy := -int(math.Abs(float64(y2 - y1)))
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx + dy
	for {
		fmt.Printf("\033[%d;%dH*", y1, x1) // Move to position and draw '*'
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x1 += sx
		}
		if e2 <= dx {
			err += dx
			y1 += sy
		}
	}
}

func main() {
	fmt.Print("\033[H\033[2J") // Clear the screen
	fmt.Println("Recursive Tree Visualization:")

	// Start the tree at the center of the screen
	screenWidth := 80
	screenHeight := 24
	rootX, rootY := screenWidth/2, screenHeight-1

	// Draw the tree
	drawTree(rootX, rootY, 12, -math.Pi/2, 6)

	// Keep the visualization on screen
	fmt.Print("\033[0m\nPress Enter to exit...")
	fmt.Scanln()
}
