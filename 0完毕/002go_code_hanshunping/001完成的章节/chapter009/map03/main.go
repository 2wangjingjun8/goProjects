package main

import (
	"fmt"
)

func main()  {
	var cities = make(map[string]string)
	cities["c01"] = "广州"
	cities["c03"] = "湛江"
	cities["c04"] = "深圳"
	cities["c05"] = "惠州"
	fmt.Println("删除前cities=",cities)
	// 第一种删除delete()
	delete(cities,"c04")

	fmt.Println("第一种删除后cities=",cities)

	// 第二种删除
	// cities = make(map[string]string)
	// fmt.Println("第二种删除后cities=",cities)

	// 查找
	val := cities["c05"]
	fmt.Println("val = ",val)

	// val,ok := cities["c05"]
	// fmt.Println("ok = ",ok)
	// if ok {
	// 	fmt.Println("val = ",val)
	// }else{
	// 	fmt.Println("不存在")

	// }




}