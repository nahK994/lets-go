package main

import "fmt"

func main() {
	fmt.Println("Give me two number")
	fmt.Print("first number: ")
	var a int
	fmt.Scan(&a)
	fmt.Print("second number: ")
	var b int
	fmt.Scan(&b)
	var c int
	c = a + b
	fmt.Println(a, "+", b, "=", c)
}
