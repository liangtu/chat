package main

import (
	"chatroom/client/process"
	"fmt"
	"github.com/gookit/color"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	up := &process.UserProcess{}
	for {
		fmt.Println()
		color.Info.Println("******************** 欢迎登录聊天系统 ********************")
		fmt.Println()
		color.Magenta.Println("\t\t\t1 登陆聊天系统")
		fmt.Println()
		color.Cyan.Println("\t\t\t2 注册用户")
		fmt.Println()
		color.Blue.Println("\t\t\t3 退出系统")
		fmt.Println()
		color.Info.Println("**************************************************************")
		fmt.Println()
		color.Note.Println("请选择（1-3）：")
		fmt.Println()
		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			return
		}
		switch key {
		case 1:
			color.Red.Println("请输入用户ID:")
			_, err := fmt.Scanf("%d\n", &userId)
			fmt.Println()
			if err != nil {
				return
			}
			color.Red.Println("请输入用户密码:")
			fmt.Println()
			_, err = fmt.Scanf("%s\n", &userPwd)
			if err != nil {
				return
			}
			err = up.Login(userId, userPwd)
			if err != nil {
				return
			}

		case 2:
			color.Red.Println("请输入用户ID")
			_, err2 := fmt.Scanf("%d\n", &userId)
			if err2 != nil {
				return
			}

			color.Red.Println("请输入用户密码")
			_, err := fmt.Scanf("%s\n", &userPwd)
			if err != nil {
				return
			}
			color.Red.Println("请输入昵称")
			_, err = fmt.Scanf("%s\n", &userName)
			if err != nil {
				return
			}

			err = up.Register(userId, userPwd, userName)
			if err != nil {
				return
			}
			os.Exit(0)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误")
		}
	}
	return
}
