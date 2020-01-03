package main

import(
	"fmt"
	"net"
	"go_code/chatroom/message"
	"encoding/json"
	_"errors"
	"io"
)

func serverProcesLogin(conn net.Conn, mes message.Message ) (err error)  {
	// 1.直接从mes中取出mes.Data，并直接反序列化为LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
	}
	// 2. 声明loginResMes,并完成赋值
	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	}else{
		loginResMes.Code = 500
		loginResMes.Error = "账号密码错误..."
	}
	// 3.loginResMes 序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}
	// 4.声明resMes,将data赋值给resMes.Data
	var resMes message.Message
	resMes.Type = "LoginResMesType"
	resMes.Data = string(data)
	// 5.将resData序列化，发送
	data,err = json.Marshal(resMes) 
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
	}
	writePkg(conn, data)
	return
}

// 根据客户端发送消息不同，决定调用哪个函数来处理
func serverProcesMes(conn net.Conn, mes message.Message ) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			// 处理登录数据消息
			serverProcesLogin(conn, mes)
		case message.RegisterMesType:
			// 处理注册数据消息
		default:
			fmt.Println("消息类型不存在，无法法处理数据...")
	}
	return
}

func process(conn net.Conn )  {
	defer conn.Close()
	// 循环接受客户发来的消息
	for{
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出。。。")
				
			}else{
				fmt.Println("err = ",err)
			}
			return
		}
		fmt.Println("mes = ",mes)
		serverProcesMes(conn, mes)

		
	}

}

func main()  {
	fmt.Println("服务器在8889端口监听...")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err = ",err)
		return
	}
	for{
		fmt.Println("正在等待客户来链接服务器...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ",err)
		}
		go process(conn)
	}
}