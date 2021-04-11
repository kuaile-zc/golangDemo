package main

import "fmt"

func main() {
	//数组初始化四种方法
	var array01 [3]int = [3]int{1, 2, 3}
	fmt.Println("Array01=", array01)

	var array02 = [3]int{5, 6, 7}
	fmt.Println("Array02=", array02)

	var array03 = [...]int{8, 9, 10}
	fmt.Println("Array03=", array03)

	var array04 = [...]int{1: 100, 0: 300, 2: 200}
	fmt.Println("Array04=", array04)

	for index, value := range array04 {
		fmt.Printf("This index is %v and the value is %v \n", index, value)
	}

}
