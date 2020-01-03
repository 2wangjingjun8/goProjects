package main

import (
	"fmt"
	"strconv"
)

func toBin(num int) string {
	bin := ""
	for ; num > 0; num /= 2 {
		res := num % 2
		bin = strconv.Itoa(res) + bin
	}
	return bin
}

func main() {
	fmt.Println(toBin(10)) //1010
}
