// Package pe implements commonly used functions for solving Project Euler problems
package pe

import "math"

// IsPrime tests primality
func IsPrime(n int) bool {

	switch {
	case n == 1:
		return false
	case n < 4:
		return true
	case n%2 == 0:
		return false
	case n < 9:
		return true
	case n%3 == 0:
		return false
	default:
		r := int(math.Floor(math.Sqrt(float64(n))))
		f := 5
		for f <= r {
			switch {
			case n%f == 0:
				return false
			case n%(f+2) == 0:
				return false
			default:
				f += 6
			}
		}
		return true
	}
}

// SieveOfEratosthenes implements the sieve of Eratosthenes
func SieveOfEratosthenes(n int) []bool {
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false

	for i, v := range sieve {
		if v {
			for j := i + i; j < n+1; j += i {
				sieve[j] = false
			}
		}
	}

	return sieve
}
