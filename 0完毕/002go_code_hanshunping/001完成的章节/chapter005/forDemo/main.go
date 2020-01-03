package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 1; i<=10;i++{
		go fmt.Println("hello xiaoxiao",i)
		fmt.Println("hello xiaoxiao",i)
		time.Sleep(2 * time.Second)
	}
}