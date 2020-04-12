package main

import "fmt"

func divisibleBy3(x int) bool {
	return x%3 == 0
}

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
