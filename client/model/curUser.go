package model

import (
	"chatroom/common/message"
	"net"
)

//在客户端 作为全局
type CurUser struct {
	Conn net.Conn
	message.User
	ToUserId int
}
