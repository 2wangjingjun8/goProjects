package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filePath := "aa1.txt"
	write(filePath)
}
func read(filePath string) {
	file1, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file1 err = ", err)
	}
	reader := bufio.NewReader(file1)
	// data, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("reader ReadString err = ", err)
	// }
	// fmt.Println("data = ", data)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕 ")
				break
			} else {
				fmt.Println("reader ReadString err = ", err)
				return
			}
		}
		fmt.Println("data = ", data)
	}

}

func write(filePath string) {

	file2, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file2 err = ", err)
	}
	writer := bufio.NewWriter(file2)
	for i := 0; i <= 1000; i++ {
		writer.WriteString("hello world" + strconv.Itoa(i) + "\n")
	}
	writer.Flush()
}
