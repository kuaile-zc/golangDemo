package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error occurred: %s\n", err.Error())
		}else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error."))
	b := 0
	a := 5/b
	fmt.Println(a)
}
func main() {
	tryRecover()
}
