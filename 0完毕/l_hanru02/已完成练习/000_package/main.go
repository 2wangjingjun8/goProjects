package main

import (
	"fmt"
	"l_package/utils"
	"l_package/utils/utilsTime"
)

func main() {
	utilsTime.PrintTime()
	utils.Count()
}

func init() {
	fmt.Println("123")
}
func init() {
	fmt.Println("456")
}
