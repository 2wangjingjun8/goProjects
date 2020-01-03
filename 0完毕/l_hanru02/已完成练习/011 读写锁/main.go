package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg      *sync.WaitGroup
	rwMutex *sync.RWMutex
)

func main() {
	wg = new(sync.WaitGroup)
	rwMutex = new(sync.RWMutex)
	wg.Add(2)
	// go readFile(1)
	// go readFile(2)
	go readFile(3)
	go writeFile(1)
	// go writeFile(2)
	wg.Wait()

}
func readFile(i int) {
	defer wg.Done()
	fmt.Println(i, "准备读数据...")
	rwMutex.RLock()
	fmt.Println(i, "正在读数据...")
	time.Sleep(2 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(i, "读数据结束...")
}
func writeFile(i int) {
	defer wg.Done()
	fmt.Println(i, "\t准备写数据...")
	rwMutex.Lock()
	fmt.Println(i, "\t正在写数据...")
	time.Sleep(2 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "\t写数据结束...")
}
