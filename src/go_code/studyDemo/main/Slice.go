package main

import "fmt"

func main() {
	//演示切片
	//第一种方式
	var intArr [5]int = [...]int{1, 2, 3, 7, 12}

	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("The elements of the slice are ", slice)
	fmt.Println("The number of the slice ", len(slice))
	fmt.Println("The capacity of the slice ", cap(slice))

	//第二种方式

	slice02 := make([]int, 2, 4)
	slice02[0] = 1
	slice02[1] = 100
	fmt.Println("The elements of the slice02 are ", slice02)

	//第三种
	var slice03 []string = []string{"Tom", "jack", "mary"}
	fmt.Println("The elements of the slice03 are ", slice03)

	//如何扩容
	slice03 = append(slice03, "Sam")
	fmt.Println("The elements of the slice03 are ", slice03)

	slice03 = append(slice03, slice03...)
	fmt.Println("The elements of the slice03 are ", slice03)

	str := "4618965318@qq.com"
	sliceStr := str[11:]
	fmt.Println("SliceStr=", sliceStr)

	array := []byte(str)
	array[0] = 'a'
	str = string(array)
	fmt.Println("Str=", str)

	//rune 兼容中文
	array01 := []rune(str)
	array01[0] = '啊'
	str = string(array01)
	fmt.Println("Str=", str)
}
