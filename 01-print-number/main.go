package main

import "fmt"

func main() {
	fmt.Println("Give me a number. I will tell you its double")
	fmt.Print("Number: ")
	var a int
	fmt.Scan(&a)
	var b int
	b = 2 * a
	fmt.Println(2, "x", a, "=", b)
}
