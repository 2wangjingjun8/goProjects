package main

import(
	"fmt"
	_"time"
)

func setNum(numChan chan int)  {
	for i := 1; i <= 80; i++ {
		numChan<- i
	}
	close(numChan)
}

func putNum( numChan chan int, putChan chan int, exitChan chan bool)  {
	var flag bool
	for{
		num,ok := <-numChan
		if !ok {
			break
		}
		if num == 1 {
			continue
		}
		flag = true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			putChan<- num
		}
	}
	fmt.Println("有一个协程获取不了数据了，退出")
	exitChan<- true

}

func main()  {
	numChan := make(chan int, 1000)
	putChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go setNum(numChan)
	for i := 0; i < 4; i++ {
		go putNum(numChan, putChan, exitChan)
	}
	
	go func ()  {
		for i := 0; i < 4; i++ {
			<-exitChan
		}	
		close(putChan)
	}()

	for{
		v,ok := <-putChan
		if !ok {
			break
		}
		fmt.Println("素数：",v)

	}
	fmt.Println("主线程退出")

}