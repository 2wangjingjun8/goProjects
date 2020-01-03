package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	myMap = make(map[int]uint, 10)
	lock sync.Mutex
)

func test(n int)  {
	var res uint = 1
	for i := 1; i <= n; i++ {
		res *= uint(i)
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
	fmt.Printf("myMap[%d] = %d\n",n,res)
}

func main()  {

	for i := 0; i < 30; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)

}