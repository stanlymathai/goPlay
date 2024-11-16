/*
Langton's Ant is a cellular automation that shows how a simple set of rules can lead to chatotic
behavior and eventually produce surprising order.package goplay
*/

package main

import (
	"fmt"
	"time"
)

const (
	width  = 40 // Width of the grid
	height = 20 // Height of the grid
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Ant struct {
	x, y      int
	direction Direction
}

// initializeGrid creates an empty grid
func initializeGrid() [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

// displayGrid prints the grid and the position of the ant
func displayGrid(grid [][]bool, ant Ant) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == ant.y && j == ant.x {
				fmt.Print("A") // Ant position
			} else if grid[i][j] {
				fmt.Print("█") // Colored cell
			} else {
				fmt.Print(" ") // Empty cell
			}
		}
		fmt.Println()
	}
}

// moveAnt updates the ant's position and the grid based on the rules
func moveAnt(grid [][]bool, ant *Ant) {
	// Turn and flip the cell
	if grid[ant.y][ant.x] {
		ant.direction = (ant.direction + 3) % 4 // Turn left
	} else {
		ant.direction = (ant.direction + 1) % 4 // Turn right
	}
	grid[ant.y][ant.x] = !grid[ant.y][ant.x] // Flip the cell

	// Move forward
	switch ant.direction {
	case Up:
		ant.y = (ant.y - 1 + height) % height
	case Right:
		ant.x = (ant.x + 1) % width
	case Down:
		ant.y = (ant.y + 1) % height
	case Left:
		ant.x = (ant.x - 1 + width) % width
	}
}

func main() {
	grid := initializeGrid()
	ant := Ant{width / 2, height / 2, Up}

	fmt.Println("Langton's Ant Simulation\nPress Ctrl+C to stop...")
	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")
		displayGrid(grid, ant)
		moveAnt(grid, &ant)
		time.Sleep(100 * time.Millisecond) // Control the simulation speed
	}
}
