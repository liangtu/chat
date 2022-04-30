package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func init() {
	initPool("192.168.142.119:6379", 16, 0, 300*time.Second)
	initUserDao()
}

//初始化一个全局的操作数据库的实例
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	fmt.Println("服务器开始监听.....")
	listen, err := net.Listen("tcp", ":30405") //todo 配置文件

	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {

		}
	}(listen)

	for {
		fmt.Println("等待客户端连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)
		}
		//这里准备一个协程为客户端服务
		go processCom(conn)
	}

}
