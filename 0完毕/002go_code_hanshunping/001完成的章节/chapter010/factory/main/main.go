package main

import (
	"fmt"
	"go_code/chapter010/factory/model"
)

func main()  {
	// stu := model.Student{"haha",20}
	stu := model.GetStu("heh",50)
	aa := model.GetName(stu)

	fmt.Println(*stu)
	fmt.Println(aa)
}