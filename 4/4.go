package main

import (
	"fmt"
	"reflect"
)

func intToReverseSlice(n int) []int {
	var remainder int
	var sliced []int

	for i := 0; ; i++ {
		remainder = n % 10
		sliced = append(sliced, remainder)
		n -= remainder
		n /= 10
		if n == 0 {
			break
		}
	}

	return sliced
}

func reverseSlice(slice []int) {

	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func isPalindrome(n int) bool {
	a := intToReverseSlice(n)
	b := make([]int, len(a))
	copy(b, a)
	reverseSlice(b)
	if reflect.DeepEqual(a, b) {
		return true
	}
	return false
}

func max(slice []int) int {
	max := slice[0]

	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {

	var palindromes []int
	var n int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			n = i * j
			if isPalindrome(n) {
				palindromes = append(palindromes, n)
			}
		}
	}
	fmt.Println(max(palindromes))
}
