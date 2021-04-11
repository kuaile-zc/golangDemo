package main

import "fmt"

func main() {

	var name string
	var age int
	var sal float64

	fmt.Println("Please input name.")
	fmt.Scanln(&name)

	fmt.Println("This name is ", name)

	fmt.Println("Please input name age and salary use blank separate.")
	fmt.Scanf("%s %d %f", &name, &age, &sal)

	fmt.Printf("This person name age and salary is %s %d %f", name, age, sal)

}
