package main

import(
	"fmt"
	"io/ioutil"
)
func main() {
	content,err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("err = ",err)
	}
	fmt.Println(content)
	fmt.Println(string(content))
}