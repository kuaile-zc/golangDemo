package main

import (
	"fmt"
	"time"
)

func main() {

	//Get current time.
	now := time.Now()
	fmt.Println("New time is ", now)

	//Format time   must follow 2006-1-2 15:04:05
	fmt.Printf("%d-%d-%d %d:%d:%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	timeStr := fmt.Sprint(now.Format("2006-1-2 15:04:05"))
	fmt.Println("Time string is ", timeStr)
}
