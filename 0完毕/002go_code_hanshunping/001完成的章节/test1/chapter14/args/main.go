package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("参数长度为：",len(os.Args))
	for i,v := range os.Args{
		fmt.Printf("Args[%v] = %v\n",i,v)
	}
}