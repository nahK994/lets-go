package main

import "fmt"

func main() {
	fmt.Printf("Give me a number; ")
	var num int
	fmt.Scanf("%d", &num)
	var num2 int
	num2 = num * 2
	fmt.Printf("double of %d is %d\n", num, num2)

}
