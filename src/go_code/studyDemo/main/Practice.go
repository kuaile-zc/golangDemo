package main

import "fmt"

func main() {
	//97 have week and day
	var days int = 97
	var weeks int = days / 7
	var remainDays int = days % 7
	fmt.Printf("This 97 have %v weeks and %v days. \n", weeks, remainDays)

	//Change a,b not use temp
	a := 1
	b := 9
	fmt.Printf("old a=%v b=%v \n", a, b)
	a = b + a
	b = a - b
	a = a - b
	fmt.Printf("new a=%v b=%v \n", a, b)

}
