package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//....
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes-json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes-json.Marshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}

func (this *SmsProcess) SendSingleMes(mes *message.Message) {
	var smsMes message.SmsSingleMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendSingleMes-json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendSingleMes-json.Marshal err=", err)
		return
	}
	//用toUserId 找出 在线的用户
	up := userMgr.onlineUsers[smsMes.ToUserId]
	this.SendMesToEachOnlineUser(data, up.Conn)
}
