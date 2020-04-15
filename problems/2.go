package problems

import (
	"fmt"
	"time"
)

// Problem2 solves problem 2
func Problem2() {
	start := time.Now()

	var evenFib int

	var a int = 1
	var b int = 2

	for b < 4000000 {
		if a%2 == 0 {
			evenFib += a
		}
		a, b = b, a+b
	}
	evenFib += a

	elapsed := time.Since(start)

	fmt.Println(evenFib)
	fmt.Println(elapsed)
}
