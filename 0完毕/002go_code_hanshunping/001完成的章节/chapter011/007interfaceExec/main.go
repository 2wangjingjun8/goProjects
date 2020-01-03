package main

import (
	"fmt"
	"sort"
	"math/rand"
)

type Student struct{
	Name string
	Age int
}

type stuSlice []Student
func (stu stuSlice) Len() int {
	return len(stu)
}
func (stu stuSlice) Less(i, j int) bool {
	return stu[i].Age < stu[j].Age
}
func (stu stuSlice) Swap(i, j int) {
	stu[i],stu[j] = stu[j],stu[i]
}

func main()  {
	intSlice := []int{1,-5,3,50,65,12}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var stus stuSlice

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:fmt.Sprintf("学生——%v",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		stus = append(stus,stu)
	}

	fmt.Println("排序前:")
	for i := 0; i < len(stus); i++ {
		fmt.Println(stus[i])
	}
	fmt.Println("排序后:")
	sort.Sort(stus)
	for i := 0; i < len(stus); i++ {
		fmt.Println(stus[i])
	}


}