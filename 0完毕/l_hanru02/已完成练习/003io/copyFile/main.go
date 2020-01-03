package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "../static/go.jpg"
	destFile := "../static/go3.jpg"
	copyFile3(srcFile, destFile)
}

func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open file1 err = ", err)
	}
	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file2 err = ", err)
	}
	defer file1.Close()
	defer file2.Close()
	total := 0
	n2 := -1
	data := make([]byte, 1024, 1024)
	for {
		n1, err := file1.Read(data)
		if err == io.EOF {
			fmt.Println("file1 copy finish")
			break
		} else if err != nil {
			fmt.Println("file read err = ", err)
			return total, err
		}
		n2, err = file2.Write(data[:n1])
		if n1 != n2 {
			fmt.Println("read n1 != write n2")
		}
		total += n2
	}
	return total, nil
}

func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open file1 err = ", err)
	}
	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file2 err = ", err)
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

func copyFile3(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println("err = ", err)
		return -1, err
	}
	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("err = ", err)
		return -1, err
	}
	return len(input), nil
}
