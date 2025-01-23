package main

import "fmt"

func main() {
	fmt.Println("Give me your year of birth. I will give you your age")
	fmt.Print("Number: ")
	var a int
	fmt.Scan(&a)
	var b int
	b = 2025 - a
	fmt.Println("You are ", b, " years of old")
}
