package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	client := http.Client{}
	request, _ := http.NewRequest("GET", "https://www.zhenai.com/zhenghun", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: status code is: ",response.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printCityList(all)
}

func printCityList(contents []byte) {

	compile := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+" data-v-1573aa7c>[^<]+</a>`)
	matches := compile.FindAll(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}
}
