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

	fmt.Println(students)
	fmt.Println(students["stu01"]["addr"])
	fmt.Println(students["stu02"]["age"])
}