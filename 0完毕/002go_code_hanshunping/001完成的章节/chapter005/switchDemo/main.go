package main

import "fmt"

func test(char byte) (byte) {
	return char + 1
}

func main()  {
	// var key byte
	// fmt.Println("请输入一个字符：a,b,c,d,e")
	// fmt.Scanf("%c",&key)
	// switch key {
	// 	case 'a':
	// 		fmt.Println("今天是星期一")
	// 	case 'b':
	// 		fmt.Println("今天是星期二")
	// 	case 'c':
	// 		fmt.Println("今天是星期三")
	// 	default:
	// 		fmt.Println("输入有误")
	// }

	// switch test(key) {
	// 	case 'a':
	// 		fmt.Println("今天是星期一")
	// 	case 'b':
	// 		fmt.Println("今天是星期二")
	// 	case 'c':
	// 		fmt.Println("今天是星期三")
	// 	default:
	// 		fmt.Println("输入有误")
	// }
	var n int32
	var m int32 = 5
	fmt.Scanf("%v",&n)
	switch n {
	case 10,15,5:
		fmt.Println("今天是星期一")
	case m:
		fmt.Println("今天是星期二")
	case 'c':
		fmt.Println("今天是星期三")
	default:
		fmt.Println("输入有误")
	}

}