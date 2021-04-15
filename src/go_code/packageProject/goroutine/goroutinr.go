package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello from goroutine ", i)
			}
		}(i)
	}

	time.Sleep(time.Minute)
	//fmt.Println(a)
}
