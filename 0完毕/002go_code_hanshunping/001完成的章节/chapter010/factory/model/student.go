package model

type student struct {
	Name string
	Age int
}

func GetStu(name string,age int) *student {
	return &student{
		Name:name,
		Age:age,
	}
}


func GetName(stu *student) string {
	return stu.Name
}
