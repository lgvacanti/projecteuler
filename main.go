package main

import (
	"fmt"
	"time"

	"./problems"
)

func main() {
	start := time.Now()

	problems.Problem97()

	elapsed := time.Since(start)
	fmt.Println("Took %s", elapsed)
}
