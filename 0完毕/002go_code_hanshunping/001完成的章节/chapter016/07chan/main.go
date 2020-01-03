package main

import(
	"fmt"
	_"time"
)

func setNum(numChan chan int)  {
	for i := 1; i <= 10; i++ {
		numChan<-i
	}
	close(numChan)
}

func setSum(numChan chan int, resChan chan int, exitChan chan bool)  {
	var sum int
	for{
		v,ok := <-numChan
		if !ok {
			break
		}
		for i := 1; i <= v; i++ {
			sum += i
		}
		resChan <- sum
	}
	fmt.Println("有一个协程退出")
	exitChan <- true
}

func main()  {
	numChan := make(chan int, 0)
	resChan := make(chan int, 0)
	exitChan := make(chan bool, 8)

	go setNum(numChan)
	for i := 0; i < 8; i++ {
		go setSum(numChan, resChan, exitChan)
	}

	go func ()  {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	
	for{
		v,ok := <-resChan
		if !ok {
			break
		}
		fmt.Println("数：",v)

	}
	fmt.Println("主线程退出")
	
}