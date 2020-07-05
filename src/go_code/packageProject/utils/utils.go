package utils

import "fmt"

/**
  This is an arithmetic method.

 */
func BasicOperation(a int, b int, operation byte) float64{
	var ret float64

	switch operation {
	case '+':
		ret = float64(a) + float64(b)
	case '-':
		ret = float64(a) - float64(b)
	case '*':
		ret = float64(a) * float64(b)
	case '/':
		ret = float64(a) / float64(b)
	default:
		fmt.Println("This is a invalid operation.")
		ret = 0

	}

	return ret
}
