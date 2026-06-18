package main

import (
	rpcdemo "google/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// rpc服务器
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicf("accept error:%v", err) //accept的问题
			continue
		}
		go jsonrpc.ServeConn(conn) //处理掉
	}
}
