package main

import(
	"fmt"
	"net"
)

func Process(conn net.Conn)  {
	defer conn.Close()

	for{
		buf := make([]byte, 1024)
		// fmt.Printf("服务端正在等待客户端%s 发送信息\n",conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err = ",err)
			return
		}
		fmt.Print(string(buf[:n]))
	}

}

func main()  {
	listen,err := net.Listen("tcp",":8888") 
	if err != nil {
		fmt.Println("listen err = ",err)
	}
	for{
		fmt.Println("服务端正在监听连接...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("conn err = ",err)
			continue
		}
		fmt.Printf("conn = %v 网络地址：%v",conn,conn.RemoteAddr().String())
		go Process(conn)
	}
	
}