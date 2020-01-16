package main

import (
	"fmt"
)

func addSum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i 
	}
	return res 
}

func main()  {
	res := addSum(10)
	if res != 55 {
		fmt.Println("测试失败")
	}else{
		fmt.Printf("测试成功;结果是：%v, 期望值是：%v",res,55)
	}
}