package fibonacci

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

//斐波那契数列 1， 1， 2， 3， 5， 8， 13
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read (p []byte) (n int, err error){
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := Fibonacci()
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())

	println("----------------")
	time.Sleep(1000)

	f2 := Fibonacci()
	printFileContents(f2)
}
