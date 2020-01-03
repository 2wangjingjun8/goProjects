package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var i int8 = -125
	// var i = -125
	fmt.Println(i)
	fmt.Printf("i 的类型是%T,i的字节大小是%d",i,unsafe.Sizeof(i))
}