package main

import (
	"fmt"
	"net"
	"go_code/chatroom/message"
	"encoding/json"
)

func login(userId int, userPwd string) (err error) {
	// fmt.Printf("用户id:%v 密码：%v",userId,userPwd)
	// return nil

	// 1.链接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ",err)
		return
	}
	defer conn.Close()

	// 2.准备通过conn发送消息到服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建LoginMes消息结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd

	// 4.将LoginMes序列化
	data,err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}

	// 5.将data赋给mes.Data
	mes.Data = string(data)

	// 6.将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ",err)
		return
	}
	// 到这，data就是我们要发送的数据
	writePkg(conn,data)

	// 这里还需处理服务端返回来的消息
	mes,err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn)",err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal",err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	}else{
		fmt.Println("登录失败")
	}

	return

}