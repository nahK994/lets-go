package main

import "fmt"

func main() {
	var name string
	fmt.Print("Tell me your name: ")
	fmt.Scan(&name)
	fmt.Print("Hello, ", name, "\n")
}
