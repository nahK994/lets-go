package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var isPrime bool = true
	if n < 2 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 {
		isPrime = false
	} else {
		for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
			if n%i == 0 {
				isPrime = false
			}
		}
	}

	if isPrime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}
