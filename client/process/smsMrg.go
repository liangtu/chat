package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes-json.Unmarshal err", err)
		return
	}
	info := fmt.Sprintf("【用户id:%d 对大家说】：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}

func outputSingleMesMes(mes *message.Message) {
	var smsMes message.SmsSingleMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes-json.Unmarshal err", err)
		return
	}
	//CurUser.ToUserId = smsMes.UserId
	info := fmt.Sprintf("【用户id:%d 对你说】：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
