package main

import "fmt"

func main() {
	_, b, c := getNum()
	fmt.Println(b, c)

}

func getNum() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
