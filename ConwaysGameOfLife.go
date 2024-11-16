/*
Conway's Game of life, a cellular automation that reveals how
complexity can emerge from simple rules. It showcase the emergent
behavior and raises questions about life, order, and chaos.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 40 // Width of the grid
	height = 20 // Height of the grid
)

// initializeGrid creates a random initial grid
func initializeGrid() [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) == 1 // Randomly populate cells
		}
	}
	return grid
}

// displayGrid prints the grid to the console
func displayGrid(grid [][]bool) {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Print("█") // Alive cell
			} else {
				fmt.Print(" ") // Dead cell
			}
		}
		fmt.Println()
	}
}

// countNeighbors counts the alive neighbors of a cell
func countNeighbors(grid [][]bool, x, y int) int {
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	count := 0
	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < width && ny >= 0 && ny < height && grid[ny][nx] {
			count++
		}
	}
	return count
}

// nextGeneration computes the next state of the grid
func nextGeneration(grid [][]bool) [][]bool {
	newGrid := make([][]bool, height)
	for i := range newGrid {
		newGrid[i] = make([]bool, width)
		for j := range newGrid[i] {
			neighbors := countNeighbors(grid, j, i)
			if grid[i][j] {
				newGrid[i][j] = neighbors == 2 || neighbors == 3 // Survive
			} else {
				newGrid[i][j] = neighbors == 3 // Reproduction
			}
		}
	}
	return newGrid
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize and display the grid
	grid := initializeGrid()

	fmt.Println("Conway's Game of Life\nPress Ctrl+C to stop...")
	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")
		displayGrid(grid)
		grid = nextGeneration(grid)
		time.Sleep(200 * time.Millisecond) // Control the simulation speed
	}
}
