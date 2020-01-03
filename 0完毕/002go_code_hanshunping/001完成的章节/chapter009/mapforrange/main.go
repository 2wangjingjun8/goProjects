package main

import (
	"fmt"
)

func main()  {
	
	var students = make(map[string]map[string]string)
	students["stu01"] = make(map[string]string)
	students["stu01"]["name"] = "xiaoxiao"
	students["stu01"]["age"] = "18"
	students["stu01"]["addr"] = "广州"
	
	students["stu02"] = make(map[string]string)
	students["stu02"]["name"] = "maoamo"
	students["stu02"]["age"] = "20"
	students["stu02"]["addr"] = "湛江"

	for i,v := range students{
		fmt.Printf("i = %v, v = %v\n",i,v)
		for i1,v1 := range v{
			fmt.Printf("\t i1 = %v, v1 = %v\n",i1,v1)
		}

	}
}