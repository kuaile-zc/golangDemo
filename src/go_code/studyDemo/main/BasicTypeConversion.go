package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int32 = 100
	//Change the int 32 to float32

	var n1 float32 = float32(i)

	var i2 int8 = int8(i)

	var i3 int64 = int64(i)

	fmt.Printf("i=%v  n1=%v  i2=%v  i3=%v", i, n1, i2, i3)

	fmt.Println()
	fmt.Printf(" i type is %T", i)
	fmt.Println()

	var num1 int = 123

	var num2 float32 = 12.89

	var b bool = true

	var myChar byte = 'a'

	var str string

	//The first way that other type is converted to string.
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T str=%v\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T str=%v\n", str, str)

	//The second way that other type is converted to string.
	var num3 int = 412
	var num4 float32 = 1.32
	var b2 bool = false

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type is %T str=%v\n", str, str)

	str = strconv.FormatFloat(float64(num4), 'f', 10, 64)
	fmt.Printf("str type is %T str=%v\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type is %T str=%v\n", str, str)

	//The string converted other type.
	var strBool string = "true"

	var b3 bool

	var strInt string = "1231"

	var i4 int64

	var strFloat string = "1.91"

	var f1 float64

	i4, _ = strconv.ParseInt(strInt, 10, 64)
	fmt.Printf("i4 type %T i4=%v \n", i4, i4)

	b3, _ = strconv.ParseBool(strBool)
	fmt.Printf("b3 type %T b3=%v \n", b3, b3)

	f1, _ = strconv.ParseFloat(strFloat, 64)
	fmt.Printf("f1 type %T f1=%v \n", f1, f1)

}
