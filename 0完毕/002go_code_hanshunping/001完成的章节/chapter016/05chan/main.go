package main

import(
	"fmt"
	_"time"
)

func writeData(chan1 chan int)  {
	for i := 0; i < 50; i++ {
		chan1<- i
		fmt.Println("写数据 i = ", i)
	}
	close(chan1)
}

func readData(chan1 chan int, chan2 chan bool)  {
	for{
		v,ok := <-chan1
		if !ok {
			chan2 <- true
			close(chan2)
			break
		}
		fmt.Println("读取数据 v = ", v)
	}
	
}

func main(){
	chan1 := make(chan int, 50)
	chan2 := make(chan bool, 50)
	go writeData(chan1)
	go readData(chan1,chan2)
	for{
		_,ok := <-chan1
		fmt.Println(ok)
		if !ok {
			break
		}
	}
	// time.Sleep(time.Second*5)
}