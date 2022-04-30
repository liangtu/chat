package main

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/utils"
	"fmt"
	"io"
	"net"
)

/*
* 处理通信
 */
func processCom(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	for {
		tf := &utils.Transfer{
			Conn: conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
			} else {
				fmt.Println("readPkg err=", err)
			}
			return
		}
		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

/**
*处理消息类型
 */
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	fmt.Println(mes.Type)
	switch mes.Type {
	case message.LoginMesType:
		sp := &process.UserProcess{
			Conn: conn,
		}
		err = sp.ServerProcessLogin(mes)
	case message.RegisterMesType:
		sp := &process.UserProcess{
			Conn: conn,
		}
		//处理注册
		err = sp.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SmsSingleMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendSingleMes(mes)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}
