package problems

import "fmt"

// Problem14 solves problem 14
func Problem14() {

	var maxChain int
	var maxChainNum int
	var n int

	for i := 1; i < 1000000; i++ {

		n = i
		for length := 1; n != 1; length++ {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
			if length > maxChain {
				maxChain = length
				maxChainNum = i
			}
		}

	}

	fmt.Println(maxChainNum, maxChain)
}
