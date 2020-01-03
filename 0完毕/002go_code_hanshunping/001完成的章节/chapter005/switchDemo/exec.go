package main

import "fmt"

func main()  {
	// var char byte
	// fmt.Println("请输入a|b|c|d")
	// fmt.Scanf("%c",&char)
	// switch char {
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")
	// 	case 'c':
	// 		fmt.Println("C")
	// 	case 'd':
	// 		fmt.Println("D")
	// 	default:
	// 		fmt.Println("other···")
	// }

	var month int32
	fmt.Println("请输入月份：")
	fmt.Scanf("%v",&month)
	switch month {
		case 3,4,5:
			fmt.Println("sprint")
		case 6,7,8:
			fmt.Println("summer")
		case 9,10,11:
			fmt.Println("autumn")
		case 12,1,2:
			fmt.Println("winter")
		default:
			fmt.Println("输入有误")
	}

}