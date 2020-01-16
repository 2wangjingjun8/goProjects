package main

import(
	"fmt"
	"os"
	"bufio"
	"io"
)
func main() {
	file,err := os.Open("./test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF{
			break
		}
	}
	fmt.Println("\n文件读取结束。。。")
	
}