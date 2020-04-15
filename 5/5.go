package main

import "fmt"

func main() {
	n := 19
	condition := false

	for !condition {
		n += 19
		condition = n%20 == 0 && n%19 == 0 && n%18 == 0 && n%17 == 0 && n%16 == 0 && n%15 == 0 && n%14 == 0 && n%13 == 0 && n%12 == 0 && n%11 == 0
	}
	fmt.Println(n)
}
