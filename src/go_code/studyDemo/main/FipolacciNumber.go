package main

import "fmt"

//1,1,2,3,5,8,-----
func main() {
	PrintFipolacci(10)
}

func PrintFipolacci(n int) {

	for i := 1; i < n+1; i++ {
		fmt.Print("", Fipolacci(i))
		if i != n {
			fmt.Print(",")
		}
	}

}

func Fipolacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fipolacci(n-1) + Fipolacci(n-2)
}
