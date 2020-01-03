package main

import (
	"fmt"
	"math/rand"
	"time"
)

var socket = 10

func main() {
	// a := 1
	// go func() {
	// 	a = 2
	// 	fmt.Println("go a = ", a)
	// }()
	// a = 3
	// time.Sleep(1 * time.Second)
	// fmt.Println("main a = ", a)
	go sellSocket("售票口1")
	go sellSocket("售票口2")
	go sellSocket("售票口3")
	go sellSocket("售票口4")
	time.Sleep(3 * time.Second)
}

func sellSocket(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if socket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			socket--
			fmt.Printf("%v 售出了一张票，剩余%v张票\n", name, socket)
		} else {
			break
		}
	}
}
