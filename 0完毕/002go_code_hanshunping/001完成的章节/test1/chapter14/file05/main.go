package main

import(
	"fmt"
	"os"
	"bufio"
	"io"
)
func main() {
	file,err := os.OpenFile("./test1.txt",os.O_RDWR|os.O_APPEND ,0644)
	defer file.Close()
	if err != nil{
		fmt.Println("err = ",err)
	}
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			break
		}
		
	}

	writer := bufio.NewWriter(file)
	for i:=0; i<5;i++ {
		writer.WriteString("I love English\r\n")
	}
	writer.Flush()
	fmt.Println("ok")
}