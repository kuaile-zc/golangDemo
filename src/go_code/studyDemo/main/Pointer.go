package main

import "fmt"

func main() {
	//TBasic data structure layout in memory.
	var i int = 10

	fmt.Println("i 的地址=", &i)

	//1. ptr 是一个指针变量
	//2. ptr 的类型是 *int
	//3. ptr 本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v \n", ptr)
	fmt.Printf("ptr 的地址:%v \n", &ptr)
	fmt.Printf("ptr 指向的值:%v \n", *ptr)
}
