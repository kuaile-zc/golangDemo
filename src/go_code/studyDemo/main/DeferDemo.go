package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	//当执行到defer的时候暂时不执行，压入一个独立的栈
	//等到函数执行完成之后才会按照先入后出的方式执行。
	defer fmt.Println("Defer print 1")
	defer fmt.Println("Defer print 2")

	ret := n1 + n2
	fmt.Println("This is sum print sum ", ret)
	return ret
}

func main() {
	result := sum(10, 20)
	fmt.Println("This is main print sum ", result)
}
