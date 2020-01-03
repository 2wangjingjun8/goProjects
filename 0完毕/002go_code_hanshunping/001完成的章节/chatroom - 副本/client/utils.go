package main
import(
	"fmt"
	"net"
	"go_code/chatroom/message"
	"encoding/binary"
	"encoding/json"
)
func readPkg(conn net.Conn ) (mes message.Message,err error)  {
	var byf = make([]byte,8096)
	fmt.Println("读取服务端发送的数据...")
	
	_,err = conn.Read(byf[0:4])
	if err != nil {
		return
	}
	fmt.Println("读取到数据：",byf[0:4])

	// 切片长度还原uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(byf[0:4])

	//接受发送过来的消息
	n,err := conn.Read(byf[0:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	err = json.Unmarshal(byf[0:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) ( err error) {
	// 先把data的长度发送给服务器
	// 获取data的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var byt [4]byte
	binary.BigEndian.PutUint32(byt[0:4],pkgLen)
	// 发送长度
	n,err := conn.Write(byt[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err = ",err)
		return
	}
	
	// 发送数据
	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err = ",err)
		return
	}
	return
}