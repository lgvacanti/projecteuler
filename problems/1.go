package problems

import "fmt"

func divisibleBy3(x int) bool {
	return x%3 == 0
}

// Problem1 solves problem 1
func Problem1() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
