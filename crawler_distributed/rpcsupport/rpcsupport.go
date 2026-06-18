// 封装服务端监听和客户端连接逻辑
package rpcsupport

import (
	"log"
	"net"             //网络接口
	"net/rpc"         //Go标准RPC框架
	"net/rpc/jsonrpc" //JSON-RPC编解码器
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)                    //将service注册到RPC，使其方法可被远程调用
	listener, err := net.Listen("tcp", host) //在指定地址上监听 TCP 连接
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept() //等待客户端连接
		if err != nil {
			log.Panicf("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn) //JSON-RPC处理该连接
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) { //连接到指定 host
	conn, err := net.Dial("tcp", host) //通过TCP拨号连接服务端
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil //创建客户端并返回
}
