package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score > 80:
		g = "A"
	}

	return g
}

func coverToBin(num int) string {
	ret := ""
	for ; num > 0; num /= 2 {
		ret = strconv.Itoa(num%2) + ret
	}
	return ret
}

func chinese() {
	s := "aa中文测试"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d  %X)", i, ch)
	}

	fmt.Println()
	println("count is ", utf8.RuneCountInString(s))

	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %S)", i, ch)
	}
}

func main() {
	fmt.Println("aa")

	enums()
	fmt.Println(grade(88))
	//fmt.Println(grade(-1))
	fmt.Println(coverToBin(9)) //1001

	arr1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("[0,5] ", arr1[0:5])
	fmt.Println("[0,6] ", arr1[0:6])
	//fmt.Println("[0,7] ", arr1[0:7])
	chinese()
}
