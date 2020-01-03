package main

import (
	"fmt"
	"strings"
)

func tpl() func(int) int {
	n := 10
	return func (x int) int {
		n += x
		return n

	}
}

func getFix( fix string) func(string) string {
	return func (path string) string {
		if !strings.HasSuffix(path,fix){
			path += fix
		}
		return path
	}
}

func main()  {
	// a := tpl()
	
	// fmt.Println(a(10))
	// fmt.Println(a(10))
	// fmt.Println(a(10))

	// b := getFix(".jpg")
	
	// fmt.Println(b("10"))
	// fmt.Println(b("10.jpg"))
	// var i int
	// for i = 0; i<10; i++{
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)
}