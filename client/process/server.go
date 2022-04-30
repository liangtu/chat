package process

import (
	"chatroom/common/message"
	"chatroom/utils"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"net"
	"os"
)

//显示登陆后的界面

func ShowMenu() {
	for {
		fmt.Println()
		color.Info.Println("******************** Welcome ********************")
		fmt.Println()
		color.Cyan.Println("\t\t1. 显示在线用户列表")
		fmt.Println()
		color.Cyan.Println("\t\t2. 私聊")
		fmt.Println()
		color.Cyan.Println("\t\t3. 群聊")
		fmt.Println()
		color.Cyan.Println("\t\t4. 信息列表") //todo
		fmt.Println()
		color.Cyan.Println("\t\t5. 离线留言") //todo
		fmt.Println()
		color.Cyan.Println("\t\t6. 退出系统")
		fmt.Println()
		color.Info.Println("**************************************************")
		fmt.Println()

		var key int
		var ToUserId int
		var content string

		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			return
		}

		smsProcess := &SmsProcess{}
		switch key {
		case 1:
			fmt.Println("当前在线人数：")
			showOnlineUser()
		//fmt.Println("在线用户列表")
		case 2:
			//私聊
			fmt.Println("请输入用户id：")
			_, err := fmt.Scanf("%d\n", &ToUserId)
			fmt.Println("请输入内容：")
			_, err = fmt.Scanf("%s\n", &content)
			if err != nil {
				return
			}
			err = smsProcess.SendSingleMes(content, ToUserId)
			if err != nil {
				return
			}

		case 3:
			fmt.Println("你可以放开大家说：")
			_, err := fmt.Scanf("%s\n", &content)
			if err != nil {
				return
			}
			err = smsProcess.SendGroupMes(content)
			if err != nil {
				return
			}
		case 4:
			fmt.Println("信息列表")
		case 5:
			fmt.Println("信息列表")
		case 6:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}

}
func SeverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println(" tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1、取出推送状态
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				return
			}

			//2、把这个用户信息保存在客户map[int]User
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		case message.SmsSingleMesType:
			outputSingleMesMes(&mes)

		default:
			fmt.Println("服务器返回了我不能处理的消息类型")
		}

	}
}
