package main

import (
	"fmt"
)

func main() {
	i := 100
	j := new(int)
	*j = 200
	fmt.Printf("i的类型：%T,i的值：%v,i的地址：%v\n",i,i,&i)
	fmt.Printf("j的类型：%T,j的值：%v,j的地址：%v,j指向的值：%v\n",j,j,&j,*j)
}
