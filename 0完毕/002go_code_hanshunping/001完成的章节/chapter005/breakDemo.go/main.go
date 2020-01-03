package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	var count int= 0
	for{
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println("num = ",num)
		count++
		if num == 99 {
			break
		}
	}
	fmt.Printf("生成99，一共用了count = %d次\n",count)
	// lable1:
	for i:=0;i<10;i++{
		lable2:
		for j:=0;j<10;j++{
			if j>=2{
				// break lable1
				break lable2
			}
			fmt.Println("j = ",j)
		}
	}
	

}