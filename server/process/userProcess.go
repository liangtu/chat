package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

//通知所有在线的用户
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知每个用户
		up.NotifyOneUser(userId)
	}
}

//通知用户
//通知所有在线的用户
func (this *UserProcess) NotifyOneUser(userId int) {
	//组装消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal!!! err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal!!! err=", err)
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyOneUser err= ", err)
		return
	}

}

/*
处理登陆相关
*/
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1、从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
	}
	//先申明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//申明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//准备登陆校验
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISIT {
			loginResMes.Code = 5001
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器出现了问题"
		}
	} else {
		loginResMes.Code = 200
		//登录成功之后把userId 和UserProcess 写入map
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		this.NotifyOthersOnlineUser(loginMes.UserId)
		fmt.Println("登陆成功", user)
	}

	data, errs := json.Marshal(loginResMes)

	if errs != nil {
		fmt.Println("json.Marshal fail221", errs)
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail! = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

/*
处理注册相关
*/
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1、从mes中取出mes.Data,并直接反序列化成LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
	}
	//先申明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//准备注册
	err = model.MyUserDao.Register(&registerMes)

	//申明一个registerResMes,并完成赋值
	var registerResMes message.RegisterResMes

	if err != nil {
		if err == model.ERROR_USER_EXISIT {
			registerResMes.Code = 500
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 501
			registerResMes.Error = err.Error()
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	data, errs := json.Marshal(registerResMes)

	if errs != nil {
		fmt.Println("json.Marshal fail221", errs)
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail! = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
