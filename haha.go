package main

import "fmt"

func main() {
	var i int
	var j int
	fmt.Scan(&i)

	j = 1
	for j <= 10 {
		fmt.Print(i, "x", j, "=", i*j, "\n")
		j++
	}
}
