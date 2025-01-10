package main

import "fmt"

func main() {
	var a int
	fmt.Printf("Helo! please give me a number: ")
	fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Printf("%d is an even number\n", a)
	} else {
		fmt.Printf("%d is an odd number\n", a)
	}
}
