package main

import (
	"fmt"
)

func main()  {
	strArr := [3]string{"xiaoxiao","maomao","haah"}
	for index,value := range strArr{
		fmt.Printf("index = %v , value = %v\n",index,value)
	}
	for _,value := range strArr{
		fmt.Printf(" value = %v\n",value)
	}
	
}