package main

import "fmt"

func main() {
	fmt.Println("The 1 day has peach ", MonkeyEatPeach(1))
}

//猴子吃桃子，每天吃一半再多吃一个，第十天再想吃的时候发现只剩一个问第一天有几个桃子
//思路分析
//第十天有桃子一个
//第九天有  （1+1）*2 = 4
//第n天有  peach(n) = (peach(n+1) +1)*2
func MonkeyEatPeach(day int) int {

	if day < 1 || day > 10 {
		fmt.Println("Days are not legal.")
		return 0
	}

	if day == 10 {
		return 1
	}

	return (MonkeyEatPeach(day+1) + 1) * 2

}
