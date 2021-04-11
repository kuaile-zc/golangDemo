package main

import "fmt"

func main() {

	var num int = 0
	var sum int = 0
	var max int = 100

	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			num++
			sum += i
			fmt.Println(i)
		}
	}

	fmt.Printf("This result is num %d sum %d \n", num, sum)
}
