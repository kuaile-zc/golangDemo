package main

import (
	"bufio"
	"fmt"
	"os"
	"go_code/packageProject/functional/fibonacci"
)

func tryDefer() {
	defer fmt.Println(1)
	panic("out of program")
	defer fmt.Println(2)
	fmt.Println(3)
	return
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}


func openFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL | os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		}else {
			fmt.Printf("%s %s %s \n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	defer file.Close()

}
func main() {
	tryDefer()
	//writeFile("fibonacci.txt")
	//openFile("fibonacci.txt")
}
