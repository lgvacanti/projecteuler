package problems

import (
	"fmt"

	"../pe"
)

// Problem146 solves problem 146
func Problem146() {
	// Does not check consecutivity yet
	var sum int
	var square int
	sieve := pe.SieveOfEratosthenes(1000000*1000000 + 27)

	for i := 1; i < 1000000; i++ {
		square = i * i
		switch {
		case !sieve[square+1]:
		case !sieve[square+3]:
		case !sieve[square+7]:
		case !sieve[square+9]:
		case !sieve[square+13]:
		case !sieve[square+27]:
		default:
			sum += i
		}
	}
	fmt.Println(sum)
}
