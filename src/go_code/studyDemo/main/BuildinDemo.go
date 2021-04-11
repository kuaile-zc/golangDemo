package main

import "fmt"

func main() {

	num1 := 100
	fmt.Printf("num1的类型是%T, num1的值是%v, num1的地址是%v \n", num1, num1, &num1)

	num2 := new(int) // num2 is *int
	fmt.Printf("num2的类型是%T, num2的值是%v, num2的地址是%v, num2的值指向的地址的值是%v \n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2的类型是%T, num2的值是%v, num2的地址是%v, num2的值指向的地址的值是%v \n", num2, num2, &num2, *num2)
}
