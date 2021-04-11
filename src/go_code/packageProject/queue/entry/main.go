package main

import (
	"go_code/packageProject/queue"
)

func main() {

	q := queue.Queue{}

	println(q.IsEmpty())
	q.Push(2)
	q.Push(3)
	q.Push(4)
	println(q.Pull())
	println("a")

}
