package main

import (
	"fmt"
)

func main()  {
	var arrays [4][6]int

	arrays[1][2] = 1
	arrays[2][1] = 2
	arrays[2][3] = 3


	for i := 0; i< len(arrays); i++ {
		for j := 0; j < len(arrays[0]); j++ {
			fmt.Print(" ", arrays[i][j])
		}
		fmt.Println()
	}
}