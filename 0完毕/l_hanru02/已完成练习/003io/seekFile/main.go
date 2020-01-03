package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "../static/go.jpg"
	seekFile(filePath)
}
func seekFileTest(filePath string) {
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err = ", err)
	}
	defer fp.Close()
	data := []byte{0}
	fmt.Println("length = ", len(data))

	_, err = fp.Read(data)
	fmt.Println("data = ", string(data))

	_, err = fp.Seek(2, 0)
	_, err = fp.Read(data)
	_, err = fp.Seek(3, 1)
	_, err = fp.Read(data)

	_, err = fp.Seek(3, 2)
	_, err = fp.Read(data)

	fmt.Println("data = ", string(data))

	if err != nil {
		fmt.Println("err = ", err)
	}

}

func seekFile(filePath string) {
	tempPath := filePath[strings.LastIndex(filePath, "/")+1:] + "temp.txt"
	// fmt.Println(tempPath)
	destPath := filePath[strings.LastIndex(filePath, "/")+1:]
	file1, err := os.Open(filePath)
	handleErr(err)
	file2, err := os.OpenFile(destPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	handleErr(err)
	file3, err := os.OpenFile(tempPath, os.O_CREATE|os.O_RDWR, 0644)
	handleErr(err)
	defer file1.Close()
	defer file2.Close()

	// step1:先读取临时文件的数据
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	// handleErr(err)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	// handleErr(err)
	fmt.Println(count)

	// step2:设置读写位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n2 := -1 // 读取的数据量
	n3 := -1 // 写的数据量

	total := int(count) //读取的总量
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制成功")
			file3.Close()
			os.Remove(tempPath)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		// 将复制的总量，存储在临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Println("total = ", total)
		// if total >= 8000 {
		// 	fmt.Println("模拟断电")
		// 	break
		// }
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
