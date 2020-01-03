package main

import (
	"fmt"
	"os"
)

func main() {
	// fileInfo文件信息
	fileInfo, err := os.Stat("./aa.txt")
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Sys())
}
