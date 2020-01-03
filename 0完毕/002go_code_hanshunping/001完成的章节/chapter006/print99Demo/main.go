package main

import (
	"fmt"
)
// 1*1 = 1
// 1*2 = 2 2*2 = 4
// 1*3 = 3 2*3 = 6 3*3 = 9

func print99(num int)  {
	
	for i := 1; i <= num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", i, j, i*j )
		}
		fmt.Println("")
	}
}

func main()  {
	var num int 
	for{
		fmt.Println("请输入九九乘法表的层数：")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		print99(num)
	}
}