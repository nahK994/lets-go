package main

import "fmt"

func main() {
	var height float32 = 5
	var length float32 = 6
	var width float32 = 8.7

	var volume float32 = height * length * width
	var area float32 = length * width

	fmt.Printf("There is a rectangle with height, width and length are respectively %f, %f, %f\n", height, length, width)
	fmt.Printf("It's area = %f, volume = %f\n", area, volume)
}
