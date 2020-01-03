package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

// MathUtil is a struct
type MathUtil struct {
}

// CalCircleArea is a function
func (m *MathUtil) CalCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}
func main() {
	mathUtil := new(MathUtil)
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}
	// 通过该函数把mathUtil中的服务注册到HTTP协议上，
	// 方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
