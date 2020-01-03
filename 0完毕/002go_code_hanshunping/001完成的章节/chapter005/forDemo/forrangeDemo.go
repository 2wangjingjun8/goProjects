package main

import (
	"fmt"
)

func main()  {
	// var str string = "hello world 北京"
	// str2 := []rune(str)
	// for i := 0;i < len(str2);i++{
	// 	fmt.Printf("%c\n",str2[i])
	// }
	// for index,val := range str{
	// 	fmt.Printf("index=%d val=%c \n",index,val)
	// }

	// var count,sum int = 0,0;
	// for i := 1; i <= 100; i++{
	// 	if i%9 == 0 {
	// 		count++
	// 		sum += i
	// 		fmt.Println(i)
	// 	}
	// }

	// fmt.Printf("count=%d sum=%d \n",count,sum)

	for i,j:=0,6;i<=6;i++{
		fmt.Printf("%d + %d = %d \n",i,j,i+j)
		j--
	}
}