package problems

import (
	"fmt"
	"math/big"
)

// Problem20 solves problem 20
func Problem20() {
	fact := big.NewInt(0)
	fact = fact.MulRange(1, 100)
	facts := fact.String()

	var sum int
	for i := 0; i < len(facts); i++ {
		sum += int(facts[i]) - 48
	}
	fmt.Println(sum)
}
