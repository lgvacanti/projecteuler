package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

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
