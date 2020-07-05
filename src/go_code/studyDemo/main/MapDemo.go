package main

import (
	"fmt"
	"sort"
)

func main()  {
	var strMap map[string]string
	//一定要分配内存
	strMap = make(map[string]string, 1)
	strMap["01"] = "Sam"
	fmt.Println(strMap)

	strMap02 := make(map[string]string)
	strMap02["01"] = "Sam"
	strMap02["02"] = "Haig"
	strMap02["03"] = "Kevin"
	fmt.Println(strMap02)

	delete(strMap02,"03")
	fmt.Println("Delete kevin ,and the map is ",strMap02)

	val, isHave := strMap02["01"]
	if isHave {
		fmt.Printf("StrMap have 01 value is %v . \n",val)
	}

	//切片map
	var carBrands  []map[string]string
	carBrands = make([]map[string]string, 2)

	carBrands[0] = make(map[string]string)
	carBrands[0]["Brand"] = "BenChi"
	carBrands[0]["name"] = "S600"

	carBrands[1] = make(map[string]string)
	carBrands[1]["Brand"] = "Baoma"
	carBrands[1]["name"] = "X6"

	fmt.Println("The carBrands is ",carBrands)

	newCarBrand := map[string]string{
		"Brand":"Aodi",
		"name":"A6",
	}

	carBrands = append(carBrands, newCarBrand)

	fmt.Println("The carBrands is ",carBrands)

	intMap := make(map[int]int)
	intMap[2] = 2
	intMap[13] = 13
	intMap[1] = 1
	intMap[9] = 9

	fmt.Println("The int map is ",intMap)

	var keys  []int
	for k,_ := range intMap  {
		keys = append(keys, k)
	}

	//Sort
	sort.Ints(keys)
	fmt.Println("Keys is ",keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}
