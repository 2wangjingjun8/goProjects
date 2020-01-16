package main

import(
	"fmt"
)

func main() {
	aa := "20背景"
	// fmt.Printf("i 的类型是%T",i)
	fmt.Printf("i 的长度是%v\n",len(aa))
	b := []rune(aa)
	fmt.Println(b)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%+v:%c\n",i,b[i])
	}
}