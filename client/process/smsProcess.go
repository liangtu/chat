package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
	//......
}

//群发
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("client-json.Marshal-SendGroupMes err=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("client-json.Marshal-SendGroupMes!! err=", err)
	}

	//发送给服务器：
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg-SendGroupMes!! err=", err)
	}
	return
}

//私聊
func (this *SmsProcess) SendSingleMes(content string, touserId int) (err error) {
	var mes message.Message
	mes.Type = message.SmsSingleMesType

	var smsMes message.SmsSingleMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	smsMes.ToUserId = touserId

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("client-json.Marshal-SendGroupMes err=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("client-json.Marshal-SendGroupMes!! err=", err)
	}

	//发送给服务器：
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg-SendGroupMes!! err=", err)
	}
	return
}
