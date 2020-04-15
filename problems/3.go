package problems

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Problem3 solves problem 3
func Problem3() {

	n := 600851475143

	for !isPrime(n) {

		for i := 2; i < n; i++ {
			if n%i == 0 {
				fmt.Println(i)
				n = n / i
				break
			}

		}
	}
	fmt.Println(n)
}
