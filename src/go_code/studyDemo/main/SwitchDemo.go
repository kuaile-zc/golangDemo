package main

import (
	"fmt"
)

func main() {

	var age int
	fmt.Println("Please input age.")
	fmt.Scanln(&age)

	switch age {
	case 1, 10:
		fmt.Println("You age is 1 or 10! good.")
	case 2:
		fmt.Println("You age is 2! Yea.")
	default:
		fmt.Println("No match age!")

	}

	var score float32
	fmt.Println("Please input score.")
	fmt.Scanln(&score)

	switch {
	case score >= 90:
		fmt.Println("This is a very good score!")
	case score >= 70 && score < 90:
		fmt.Println("This is a good score!")
	case score >= 60 && score < 70:
		fmt.Println("This is a common score!")
	default:
		fmt.Println("this is a bad score!")
	}

	var num int = 10

	switch num {
	case 1, 10:
		fmt.Println("You age is 1 or 10! good.")
		fallthrough //穿透到下一层不会判断直接执行
	case 20:
		fmt.Println("You age is 20! Yea.")
	default:
		fmt.Println("No match age!")

	}

	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x type is :%T.", i)
	case int:
		fmt.Println("x type is int.")
	case float64:
		fmt.Println("x type is float64.")
	case func(int) float64:
		fmt.Println("x type is func(int).")
	case bool, string:
		fmt.Println("x type is bool or string.")
	default:
		fmt.Println("unknow.")

	}

}
