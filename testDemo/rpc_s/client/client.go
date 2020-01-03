package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	var req float32
	req = 3
	var resp *float32

	err = client.Call("MathUtil.CalCircleArea", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)
}
