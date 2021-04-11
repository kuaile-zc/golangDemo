package main

import "fmt"

func main() {
	sum := Sum(1, 2, 4, 6, 7, 8, 2, 11, 87, 44)
	//sum := Sum(1)
	fmt.Println("Sum is ", sum)
}

func Sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
