package main

import (
	"fmt"
)

func main()  {
	var monsters []map[string]string
	monsters = make([]map[string]string,2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "maomao"
		monsters[0]["age"] = "200"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "xiaoxiao"
		monsters[1]["age"] = "100"
	}

	fmt.Println(monsters)
	newmonster := map[string]string{
		"name":"dada",
		"age":"50",
	}
	monsters = append(monsters,newmonster)
	fmt.Println(monsters)
}