package main

import (
	"fmt"
)

func main()  {
	var str = "helloworld@qq.com"
	slice := str[10:]
	fmt.Println(slice)

	// slice2 := []byte(str)
	// slice2[0] = 'z'
	// str = string(slice2)
	// fmt.Println(str)

	slice2 := []rune(str)
	slice2[0] = '北'
	str = string(slice2)
	fmt.Println(str)

}