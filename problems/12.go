package problems

import (
	"fmt"
	"math"
)

func numDivisors(n int) int {
	if n == 1 {
		return 1
	}
	numDiv := 0
	sqr := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqr; i++ {
		if n%i == 0 {
			numDiv += 2
		}
	}
	if sqr*sqr == n {
		numDiv--
	}
	return numDiv
}

// Problem12 solves problem 12
func Problem12() {
	triNum := 1
	for i := 2; true; i++ {
		triNum += i
		if numDivisors(triNum) > 500 {
			break
		}
	}
	fmt.Println(triNum)
}
