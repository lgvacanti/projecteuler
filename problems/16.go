package problems

import (
	"fmt"
	"math/big"
)

// Problem16 solves problem 16
func Problem16() {
	x := big.NewInt(2)
	y := big.NewInt(1000)
	z := big.NewInt(0)
	z = z.Exp(x, y, z)
	fmt.Println(z)

	num := z.String()
	var sum int
	for i := 0; i < len(num); i++ {
		sum += int(num[i]) - 48
	}
	fmt.Println(sum)
}
