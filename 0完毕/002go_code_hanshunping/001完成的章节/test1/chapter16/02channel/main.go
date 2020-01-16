package main
import (
	"fmt"
)

type Cat struct{
	Name string
	Age int
}

func main()  {
	intChan := make(chan int, 200)
	for i := 0; i < 100; i++ {
		intChan <- i*2
	}
	fmt.Println(intChan)
	close(intChan)
	for v := range intChan{
		fmt.Println(v)
	}
}