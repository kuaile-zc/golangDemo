package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var str = "This is a string."

	//判断长度
	length := len(str)
	fmt.Println("Str lenght is ", length)

	//中文占3个字符，解决方法
	var str2 = "test中文"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \n", r[i])
	}

	//字符串转整数
	//n ,err := strconv.Atoi("12")
	n, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("Convert error :", err)
	} else {
		fmt.Println("Convert result is ", n)
	}

	//整数转字符串
	str = strconv.Itoa(123)
	fmt.Println("Str result is ", length)

	//字符串 转[]byte
	var bytes = []byte("Hello go")
	fmt.Printf("bytes is %v \n", bytes)

	//[]byte 转字符串
	str = string([]byte{72, 101, 108, 108, 111, 32, 103, 111})
	fmt.Println(str)

	//10进制转换2,8,16
	str = strconv.FormatInt(11, 2)
	fmt.Printf("Binary is %v \n", str)

	//查找字符串中是否含有某个字符串
	b := strings.Contains("This is a beautiful sea beach", "sea")
	fmt.Printf("This str has sea %v \n", b)

	//统计字符串含有几个指定的子串
	num := strings.Count("beautiful", "u")
	fmt.Printf("The beautiful has %v char u. \n", num)

	//作比较的时候 == 会比较大小写 EqualFold不区分
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("abc equal Abc is %v \n", b)
	fmt.Printf("abc == Abc is %v \n", "abc" == "Abc")

	//返回第一次出现的Index值,没有返回-1
	index := strings.Index("IN_ADB_abd", "abd")
	fmt.Printf("The frist find index is %v. \n", index)

	//将指定的子串全部替换成你想要的换的子串 -1 表示换所有 ，1表示换一个
	str = strings.Replace("Let's go, we will to use go.", "go", "golang", -1)
	fmt.Println(str)

}
