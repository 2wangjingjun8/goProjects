package main
import (
	"fmt"
)

func writeData(intChan chan int)  {
	for i := 0; i <= 50; i++ {
		intChan<-i
		fmt.Println("写进数据 i = ",i)
	}
	close(intChan)
	
}

func readData(intChan chan int,exitChan chan bool)  {
	for{
		v,ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取数据 v = ",v)
	}
	exitChan<- true
	close(exitChan)
}

func main()  {
	var intChan = make(chan int, 100)
	var exitChan = make(chan bool, 100)
	go writeData(intChan) 
	go readData(intChan, exitChan) 
	for{
		_,ok := <-exitChan
		if !ok {
			break
		}

	}
}