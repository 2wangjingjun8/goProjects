package main

import (
	"fmt"
)

func modifyMapData(map1 map[string]map[string]string, name string){
	if map1[name] != nil {
		map1[name]["pwd"] = "888888"
	}else{
		map1[name] = make(map[string]string,2)
		map1[name]["nickName"] = "昵称"+name
		map1[name]["pwd"] = "888888"
	}
}

func main()  {
	var map1 = make(map[string]map[string]string, 0)
	map1["dada"] = make(map[string]string,2)
	map1["dada"]["nickName"] = "erha"
	map1["dada"]["pwd"] = "erha"
	modifyMapData(map1,"xiaoxiao")
	modifyMapData(map1,"maomao")
	modifyMapData(map1,"dada")
	fmt.Println(map1)
}