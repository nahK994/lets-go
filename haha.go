package main

import "fmt"

func main() {
	var year int
	var isLeapYear bool
	fmt.Scan(&year)
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 != 0 {
				isLeapYear = true
			} else {
				isLeapYear = false
			}
		} else {
			isLeapYear = true
		}
	} else {
		isLeapYear = false
	}

	if isLeapYear {
		fmt.Print(year, " is a leap year\n")
	} else {
		fmt.Print(year, " isn't a leap year\n")
	}
}
