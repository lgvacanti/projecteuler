package problems

import "fmt"

// Problem6 solves problem 6
func Problem6() {
	var sumSquares int
	squareSum := 101 * 50 * 101 * 50

	for i := 1; i < 101; i++ {
		sumSquares += i * i
	}
	fmt.Println(squareSum - sumSquares)
}
