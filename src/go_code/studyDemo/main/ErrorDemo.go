package main

import (
	"errors"
	"fmt"
)

func errorTest01() {
	//使用defer + recover 处理错误
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Error ", error)
		}
	}()

	num1 := 0
	num2 := 10
	ret := num2 / num1
	fmt.Println("Ret is ", ret)
}

func test02() {
	err := readConfig("config02.ini")
	if err != nil {
		//Print error and stop code.
		panic(err)
	}

	fmt.Println("Test 02 work...")
}

func readConfig(name string) error {
	if name == "config.ini" {
		//读取正常
		return nil
	} else {
		//返回自定义错误
		return errors.New("Read file error!")
	}
}

func main() {
	errorTest01()
	fmt.Println("Complete error test.")
	test02()
}
