package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){

	var count int = 0

	for  {
		rand.Seed(time.Now().UnixNano()+ int64(count))
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Println("生成 99 一共用了 ", count)

}
