package main

import (
	"fmt"
	"go_code/studyDemo/model"
	"go_code/studyDemo/utils"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t12\t河北\t北京")

	var a int64 = 643
	fmt.Printf("a 的数据类型为%T", a)
	fmt.Println("aa")

	var name string = "您尽快"
	fmt.Println(name)
	fmt.Printf("name = %s", name)

	//反引号的位置在1数字键左边
	str3 := `
       package main

		import "fmt"
		//加入反引号
		func main()  {
			fmt.Println("Hello")
			fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t12\t河北\t北京")

			var a int64 = 643
			fmt.Printf("a 的数据类型为%T",a)
			fmt.Println("aa")
		
			var name string = "您尽快"
			fmt.Println(name)
			fmt.Printf("name = %s",name);
	}
	`
	fmt.Println(str3)

	fmt.Println(model.HeroName)

	fmt.Printf("This is an arithmetic result %.2f \n", utils.BasicOperation(10, 3, '/'))

	method := getSum
	fmt.Printf("method type is %T, getSum type is %T \n", method, getSum)
	sum := method(1, 2)
	fmt.Println("sum is ", sum)
}

func getSum(n1 int, n2 int) (n3 int) {
	n3 = n1 + n2
	return
}
