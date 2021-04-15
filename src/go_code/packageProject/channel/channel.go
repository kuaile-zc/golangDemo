package main

import "fmt"

func createWork(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c \n", id, <-c)
		}
	}()
	return c
}
func chanDemo() {
	var channels [10]chan int

		for i := 0; i < 10; i++{
			channels[i] = createWork(i)
		}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i;
	}
	for true {

	}
}
func main() {
	chanDemo()
}
