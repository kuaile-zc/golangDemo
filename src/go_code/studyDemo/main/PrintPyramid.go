package main

import "fmt"

var ASTERISK_STR = "*"
var BLANK_STR = " "

func main() {
	PrintPyramid(10)
}

func PrintPyramid(layer int) {
	var num int = 2*(layer-1) + 1
	for i := 1; i <= num; i += 2 {
		blankNum := (num - i) / 2
		PrintBlank(blankNum)
		PrintAsterisk(i)
		PrintBlank(blankNum)
		fmt.Println()
	}
}

func PrintBlank(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(BLANK_STR)
	}
}

func PrintAsterisk(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(ASTERISK_STR)
	}
}
