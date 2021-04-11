package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name    string
	Age     int
	Color   string
	Hobby   string
	Partner []string
	Food    map[string]string
}

func (cat Cat) eat() {
	fmt.Printf("The %v will eat... \n", cat.Name)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {

	var cat01 Cat
	cat01.Name = "Sam"
	cat01.Age = 26
	cat01.Color = "Yellow"
	cat01.Hobby = "Eat"
	cat01.Partner = make([]string, 10)
	cat01.Partner[0] = "Core"

	cat01.Food = make(map[string]string)
	cat01.Food["Fish"] = "Very like"
	cat01.Food["meat"] = "Like"
	cat01.Food["rice"] = "Hate"

	fmt.Println("This Cat is ", cat01)

	cat01.eat()

	//定义结构体，方式一
	var person Person
	person.Name = "Steven"
	person.Age = 25

	fmt.Println("Person is ", person)

	//方式二
	var person02 Person = Person{}
	person02.Name = "Kevin"
	person02.Age = 32

	fmt.Println("Person 02 is ", person02)

	//方式三
	var person03 *Person = new(Person)
	(*person03).Name = "Walter"
	person03.Age = 25 //这种写法不好理解，golang设计者在底层加上了指针*

	fmt.Println("Person 03 is ", *person03)

	//第四种方式
	var person04 *Person = &Person{}
	(*person04).Name = "Corey"
	(*person04).Age = 25

	fmt.Println("Person 04 is ", *person04)

	jsonStr, err := json.Marshal(person04)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}

	fmt.Println("JsonStr", string(jsonStr))

	//打印矩阵
	var mu MethodUtils
	mu.Print()

}
