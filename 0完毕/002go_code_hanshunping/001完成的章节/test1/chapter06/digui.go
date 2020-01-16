package main

import (
	"fmt"
	_ "test/utils"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	}else{
		fmt.Println("n = ", n)
	}
	
}


func main() {
	test2(4)
}