package process

import (
	"chatroom/client/config"
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	//
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1、连接服务器
	conn, err := net.Dial("tcp", config.ADDRESS)
	if err != nil {
		fmt.Println("client dial =", err)
		return
	}
	//延迟关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)
	//2.准备通过conn发送消息类型
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	//发送给server
	tf := utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg", err)
	}

	//从server接受登陆返回消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
	}

	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		for _, v := range loginResMes.UserId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//初始化online
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		//启一个协程和服务器端的通信
		go SeverProcessMes(conn)
		//则显示在客户端
		//登陆成功后的菜单。。。。
		ShowMenu()

	} else {
		println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1、连接服务器
	conn, err := net.Dial("tcp", config.ADDRESS)
	if err != nil {
		fmt.Println("client dial =", err)
		return
	}
	//延迟关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//2.准备通过conn发送消息类型
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	//4.将loginMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	//发送给server
	tf := utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg", err)
	}

	//从server接受登陆返回消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
	}

	//将mes的data部分反序列化成LoginResMes
	var registerResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code != 200 {
		println(registerResMes.Error)
	}

	return
}
