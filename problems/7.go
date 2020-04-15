package problems

import (
	"fmt"

	"../pe"
)

/*
func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
*/

// Problem7 solves problem 7
func Problem7() {
	var primeCount int
	for i := 2; primeCount < 10001; i++ {
		if pe.IsPrime(i) {
			primeCount++
			fmt.Println(primeCount, ":", i)
		}
	}
}
