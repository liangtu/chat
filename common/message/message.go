package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsSingleMesType        = "SmsSingleMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`
}

type LoginMes struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}

type LoginResMes struct {
	Code   int    `json:"code"` //200 登陆成功  500表示没有注册
	UserId []int  //所有在人的ID
	Error  string `json:"error"`
}

type RegisterMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` //200 注册成功
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}
type SmsMes struct {
	User
	Content string `json:"content"`
}

type SmsSingleMes struct {
	SmsMes
	ToUserId int
}
