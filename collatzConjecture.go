/* Collatz conjecture sequence generator ("3n+1" problem)
The collatz Conjeecture is a mathematical mystery: 
take any positive integer n, if it's even, divide it by 2; if it's odd, multiply it by 3 and add 1.
Repeat the process, and the conjecture claims that you'll always eventually reach 1.

This simple rule creates surprisingly complex patterns and invites deep contemplation about
the nature of numbers and patterns.
*/

package main

import (
	"fmt"
	"strings"
)

// collatzSequence generates the Collatz sequence for a given number
func collatzSequence(n int) []int {
	sequence := []int{n}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		sequence = append(sequence, n)
	}
	return sequence
}

// visualizeCollatz generates a simple ASCII visualization of the Collatz sequence
func visualizeCollatz(sequence []int) {
	max := 0
	for _, num := range sequence {
		if num > max {
			max = num
		}
	}

	fmt.Println("\nCollatz Conjecture Visualization:")
	for _, num := range sequence {
		bar := strings.Repeat("█", num*50/max) // Scale the bars for visualization
		fmt.Printf("%4d: %s\n", num, bar)
	}
}

func main() {
	var input int
	fmt.Println("Enter a positive integer to explore the Collatz Conjecture:")
	fmt.Scan(&input)

	if input < 1 {
		fmt.Println("Please enter a positive integer greater than 0.")
		return
	}

	sequence := collatzSequence(input)

	// Print the sequence
	fmt.Println("\nGenerated Collatz Sequence:")
	for i, num := range sequence {
		fmt.Printf("%d ", num)
		if (i+1)%10 == 0 {
			fmt.Println() // Add a newline every 10 numbers for readability
		}
	}

	// Visualize the sequence
	visualizeCollatz(sequence)

	// Reflect on the sequence
	fmt.Printf("\nThe sequence took %d steps to reach 1.\n", len(sequence)-1)
	fmt.Println("Do all numbers eventually reach 1? That's the beauty of the Collatz Conjecture—nobody knows for sure!")
}
