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
