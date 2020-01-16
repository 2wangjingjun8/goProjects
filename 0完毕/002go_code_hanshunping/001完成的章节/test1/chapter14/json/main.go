package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int
	BirthDay string
	Salary float64
	Skill string
}

func testStruct(){
	monster := Monster{
		Name:"孙悟空",
		Age:500,
		BirthDay:"2019-09-05",
		Skill:"fly",
	}
	monster_str,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json格式化错误:err",err)
	}
	fmt.Println(monster_str)
	fmt.Println(string(monster_str))

}
func testMap(){
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "xiaowang"
	a["age"] = "18"
	a["sex"] = "男"
	a["salary"] = 6500.00

	monster_str,err := json.Marshal(a)
	if err != nil {
		fmt.Println("json格式化错误:err",err)
	}
	fmt.Println(monster_str)
	fmt.Println(string(monster_str))

}

func testSlice(){
	var slice []map[string]interface{}
	map1 := make(map[string]interface{})
	map1["name"] = "xiaowang~~"
	map1["age"] = 18
	map1["sex"] = "男"
	map1["salary"] = 6500.00

	map2 := make(map[string]interface{})
	map2["name"] = "~~hah"
	map2["age"] = 18
	map2["sex"] = "男"
	map2["salary"] = 7000.00
	slice = append(slice,map1,map2)
	slice_str,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("json格式化错误:err",err)
	}
	fmt.Println(slice_str)
	fmt.Println(string(slice_str))
}

func main()  {
	// testStruct()
	testMap()
	testSlice()
}