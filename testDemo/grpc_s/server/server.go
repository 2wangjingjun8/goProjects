package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testDemo/grpc_s/message"
	"time"

	"google.golang.org/grpc"
)

// OrderServiceImpl is a struct
type OrderServiceImpl struct{}

// GetOrderInfo is a func
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequst) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"201910130001": message.OrderInfo{OrderId: "201910130001", OrderName: "上衣", OrderStatus: "已付款"},
		"201910130002": message.OrderInfo{OrderId: "201910130002", OrderName: "零食", OrderStatus: "已付款"},
		"201910130003": message.OrderInfo{OrderId: "201910130003", OrderName: "手机", OrderStatus: "待付款"},
	}
	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.Timestamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if request.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
