package main

import "fmt"

func main() {
	var i int
	i = 1
	for i <= 10000 {
		fmt.Print(i, "\n")
		i++
	}
}
