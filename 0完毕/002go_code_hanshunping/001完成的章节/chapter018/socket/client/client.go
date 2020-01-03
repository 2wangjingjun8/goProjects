package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main()  {

	conn, err := net.Dial("tcp", "192.168.233.1:8888")
	if err != nil {
		fmt.Println("client conn err = ",err)
	}
	reader := bufio.NewReader(os.Stdin)

	for{
		data,err := reader.ReadString('\n')
		// fmt.Printf("data = %v,类型%T",data ,data)
		
		data = strings.Trim(data," \r\n")
		if data == "exit" {
			fmt.Println("客户端退出")
		}
		if err != nil {
			fmt.Println("read err = ",err)
		}

		_,err = conn.Write([]byte(data+"\n"))
		if err != nil {
			fmt.Println("conn.Write err = ",err)
		}
		// fmt.Printf("客户端发送了%d字节的数据\n",n)
	}
	
	

}