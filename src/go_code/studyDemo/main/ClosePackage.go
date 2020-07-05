package main

import (
	"fmt"
	"strings"
)

//闭包，相当于java class
func AddUpper() func(n int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			strs := strings.Split(name, ".")
			if len(strs) > 1{
				name = strs[0]
			}
			return name + suffix
		}
		return name
	}
}

func main()  {
	f := AddUpper();

	fmt.Println(f(4))
	fmt.Println(f(10))


	f2 := makeSuffix(".jpg")
	fmt.Println("After file name was deal ",f2("winter"))
	fmt.Println("After file name was deal ",f2("person.avi"))
}
