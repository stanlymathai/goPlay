/*
Prime Spiral also known as the Ulam Spiral. This spiral shows how prime numbers, when plotted in a cerain pattern,
create visually intriguing clusters and diagonal lines that defy simple explanation.
*/

package main

import (
	"fmt"
	"math"
)

// isPrime checks whether a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// generateSpiral generates an NxN Ulam Spiral of primes
func generateSpiral(size int) [][]string {
	// Create an empty spiral
	spiral := make([][]string, size)
	for i := range spiral {
		spiral[i] = make([]string, size)
		for j := range spiral[i] {
			spiral[i][j] = " " // Fill with spaces
		}
	}

	x, y := size/2, size/2 // Start at the center
	num := 1
	step := 1

	for step < size {
		// Move right
		for i := 0; i < step; i++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				if isPrime(num) {
					spiral[y][x] = "●" // Mark primes with a filled circle
				} else {
					spiral[y][x] = "·" // Non-prime with a dot
				}
			}
			x++
			num++
		}

		// Move up
		for i := 0; i < step; i++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				if isPrime(num) {
					spiral[y][x] = "●"
				} else {
					spiral[y][x] = "·"
				}
			}
			y--
			num++
		}

		step++

		// Move left
		for i := 0; i < step; i++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				if isPrime(num) {
					spiral[y][x] = "●"
				} else {
					spiral[y][x] = "·"
				}
			}
			x--
			num++
		}

		// Move down
		for i := 0; i < step; i++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				if isPrime(num) {
					spiral[y][x] = "●"
				} else {
					spiral[y][x] = "·"
				}
			}
			y++
			num++
		}

		step++
	}

	return spiral
}

// printSpiral prints the spiral in the console
func printSpiral(spiral [][]string) {
	for _, row := range spiral {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func main() {
	var size int
	fmt.Println("Enter the size of the spiral (odd number recommended):")
	fmt.Scan(&size)

	if size < 1 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	if size%2 == 0 {
		fmt.Println("Size should ideally be an odd number. Adjusting size to", size+1)
		size++
	}

	spiral := generateSpiral(size)
	fmt.Println("\nPrime Spiral:")
	printSpiral(spiral)
}
