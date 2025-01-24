package main

import "fmt"

func main() {
	fmt.Println("Give me your temperature in ferenheit. I will give you it in celcius")
	fmt.Print("Temperature: ")
	var a float32
	fmt.Scan(&a)
	var b float32
	b = (a - 32) / 9
	fmt.Println("Temperature in Ferenheight:", a, "Temperature in Celcius:", b)
}
