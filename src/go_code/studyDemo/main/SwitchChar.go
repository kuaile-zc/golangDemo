package main

import "fmt"

func main() {

	var char byte
	fmt.Println("请输入一个小写单个字符转换成大写")
	fmt.Scanf("%c", &char)

	switch {
	case char >= 97 && char <= 101:
		fmt.Printf("Chnage %c to the %c", char, char-32)
	default:
		fmt.Println("We can't deal other char, only change a,b,c,d,e.")

	}

	//fmt.Printf("%d %d", 'a', 'A')
}
