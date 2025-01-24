package main

import "fmt"

func main() {
	fmt.Println("Give me your height and weight. i will give you your BMI")
	fmt.Print("Height: ")
	var a float32
	fmt.Scan(&a)
	fmt.Print("Weight: ")
	var b float32
	fmt.Scan(&b)
	var c float32
	c = b / (a * a)
	fmt.Println("Your height is:", a, ",Your weight is:", b, ",Your BMI is:", c)
}
