package main

import (
	"fmt"
	"io"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(0)
	}
}

func read(filePath string) {
	fp, err := os.Open(filePath)
	defer fp.Close()
	handleErr(err)
	data := make([]byte, 4, 4)
	for {
		n, err := fp.Read(data[:len(data)])
		if n == 0 || err == io.EOF {
			fmt.Println("\n读取完成")
			break
		}
		fmt.Printf(string(data))
	}
}

func write(filePath, data string) {
	fp, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fp.Close()
	handleErr(err)
	_, err = fp.Write([]byte(data))
	handleErr(err)
}

func main() {
	filePath := "../static/aa.txt"
	// read(filePath)
	data := "hello world"
	write(filePath, data)
}
