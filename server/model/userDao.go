package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局变量
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISIT
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Marshal___err=", err)
		return
	}
	return
}

//登陆校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}
	return
}

//注册
func (this *UserDao) Register(userInfo *message.RegisterMes) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	userId := userInfo.UserId

	_, err = redis.String(conn.Do("HGet", "users", userId))
	if err == nil {
		err = ERROR_USER_EXISIT
		return
	}
	userData, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	userData1 := string(userData)
	fmt.Println(userData1)
	_, err = conn.Do("HSet", "users", userId, userData1)
	return
}
