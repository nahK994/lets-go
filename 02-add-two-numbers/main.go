package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("Hello! Please give me 2 numbers. I will tell you their summation\n")
	fmt.Printf("First number: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Second number: ")
	fmt.Scanf("%d", &b)
	var c int = a + b

	fmt.Printf("%d + %d = %d\n", a, b, c)
}
