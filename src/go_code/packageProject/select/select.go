package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2  = generator(), generator()
	for true {
		select {
		case n := <- c1 :
			fmt.Println("Received from c1 :", n)
		case n := <- c2 :
			fmt.Println("Received from c2 :", n)
		default:
			//fmt.Println("No value received")

		}
	}

}
