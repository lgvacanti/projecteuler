package problems

import (
	"fmt"

	"../pe"
)

// Problem10 solves problem 10
func Problem10() {
	var sum int

	for i := 2; i < 2000000; i++ {
		if pe.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
