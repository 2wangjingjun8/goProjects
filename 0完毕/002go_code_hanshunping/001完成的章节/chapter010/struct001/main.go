package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

func main()  {
	var Cat1 Cat
	Cat1.Name = "小白"
	Cat1.Age = 10
	Cat1.Color = "白色"
	Cat1.Hobby = "吃鱼o"
	fmt.Println("Cat1 = ",Cat1)

	
	// fmt.Println("Cat1的名字是：",Cat1.Name)
	// fmt.Println("Cat1的年龄是：",Cat1.Age)
	// fmt.Println("Cat1的颜色是：",Cat1.Color)
	// fmt.Println("Cat1的爱好是：",Cat1.Hobby)
	
	fmt.Printf("Cat1dizhi是：%p\n",&Cat1)
	fmt.Printf("Cat1的名字dizhi是：%p\n",&Cat1.Name)
	fmt.Printf("Cat1的年龄dizhi是：%p\n",&Cat1.Age)
	fmt.Printf("Cat1的颜色dizhi是：%p\n",&Cat1.Color)
	fmt.Printf("Cat1的爱好dizhi是：%p\n",&Cat1.Hobby)
}