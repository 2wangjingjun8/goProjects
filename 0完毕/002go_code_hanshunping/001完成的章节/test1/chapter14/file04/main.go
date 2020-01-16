package main

import(
	"fmt"
	"os"
	"bufio"
)
func main() {
	file,err := os.OpenFile("./test1.txt",os.O_WRONLY|os.O_CREATE ,0644)
	defer file.Close()
	if err != nil{
		fmt.Println("err = ",err)
	}
	writer := bufio.NewWriter(file)
	for i:=0; i<5;i++ {
		writer.WriteString("hello world\r\n")
	}
	writer.Flush()
	fmt.Println("ok")
}