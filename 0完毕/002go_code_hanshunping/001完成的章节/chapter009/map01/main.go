package main

import (
	"fmt"
)

func main()  {
	// 第一种方式 先声明,后make
	var a map[string]string
	a = make(map[string]string, 0)
	a["name"] = "xiaoxiao"
	a["sex"] = "男"

	fmt.Println(a)

	// 第二种方式 直接make
	var b = make(map[string]string)
	b["name"] = "gg"
	b["age"] = "15"
	fmt.Println(b)

	// 第三种方式 直接赋值
	var c = map[string]string{
		"name":"xx",
		"age":"20",
	}
	fmt.Println(c)
}