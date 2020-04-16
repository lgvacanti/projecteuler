package problems

import (
	"fmt"

	"../pe"
)

// Problem58 solves problem 58
func Problem58() {
	numPrimes := 3
	total := 5
	side := 3
	last := 9

	for float64(numPrimes)/float64(total) >= 0.1 {

		for i := 0; i < 4; i++ {
			last += side + 1
			total++
			if pe.IsPrime(last) {
				numPrimes++
			}
		}
		side += 2

	}
	fmt.Println(side)
}
