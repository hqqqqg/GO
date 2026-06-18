// 先启动server/main.go
package main

import (
	"fmt"
	rpcdemo "google/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)

	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// fmt.Println(result, err)

	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// fmt.Println(result, err)
}
