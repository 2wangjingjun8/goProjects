package main

import (
	"fmt"
	"sort"
)

func main()  {
	var map1 = map[int]int{
		5:10,
		4:3,
		10:80,
		60:2,
	}
	var keys []int
	for i,_ := range map1{
		keys = append(keys,i)
	}
	sort.Ints(keys)

	for _,v := range keys{
		fmt.Printf("map1[%v] = %v\n",v,map1[v])
	}
}