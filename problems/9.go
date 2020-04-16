package problems

import (
	"fmt"
	"math"
)

// Problem9 solves problem 9
func Problem9() {
	var c float64
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			c = math.Sqrt(float64(a*a + b*b))
			if math.Trunc(c) == c && a+b+int(c) == 1000 {
				fmt.Println(a, b, c)
				fmt.Println(a * b * int(c))
			}
		}
	}
}
