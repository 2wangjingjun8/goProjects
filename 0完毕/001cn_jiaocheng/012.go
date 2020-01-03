package main
// import "fmt"

func main(){
	// var a int = 4
	// var b int32
	// var c float64
	// var ptr *int
	// fmt.Printf("a变量类型 %T\n",a)
	// fmt.Printf("b变量类型 %T\n",b)
	// fmt.Printf("c变量类型 %T\n",c)

	// ptr = &a
	// fmt.Println("a的值%d",a)
	// fmt.Println("*ptr的值%d",*ptr)
	var a int = 4
	var ptr *int
	ptr = &a
	println("a的地址为", &a);    // 4
	println("a的值为", a);    // 4
	println("*ptr为", *ptr);  // 4
	println("ptr为", ptr);    // 824633794744
	
}