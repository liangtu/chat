package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (u *UserMgr) AddOnlineUser(up *UserProcess) {

	u.onlineUsers[up.UserId] = up
}

func (u *UserMgr) DeleteOnlineUser(up *UserProcess) {
	delete(u.onlineUsers, up.UserId)
}

func (u *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return u.onlineUsers
}

func (u *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := u.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
