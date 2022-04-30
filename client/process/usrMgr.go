package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 1024)

var CurUser model.CurUser //用户登录成功后，完成初始化

//在客户端显示
func showOnlineUser() {
	for id, _ := range onlineUsers {
		println("用户ID：", id)
	}
}

func updateUserStatus(mes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[mes.UserId]
	if !ok {
		user = &message.User{
			UserId: mes.UserId,
		}
		fmt.Println("有新朋友上线：")
		fmt.Println("用户ID：", mes.UserId)
		fmt.Println()
	}
	user.UserStatus = mes.Status
	onlineUsers[mes.UserId] = user
}
