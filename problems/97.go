package problems

import (
	"fmt"
	"math/big"
)

// Problem97 solves problem 97
func Problem97() {
	x := big.NewInt(0)
	y := big.NewInt(28433)
	z := big.NewInt(2)
	a := big.NewInt(7830457)
	one := big.NewInt(1)

	x = x.Exp(z, a, x)
	x = x.Mul(x, y)
	x = x.Add(x, one)
	str := x.String()
	length := len(str)
	fmt.Println(str[length-10 : length])
}
