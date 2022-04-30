package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := make([]byte, 8096)
	//conn.Read 在conn没有被关闭的的情况下才会阻塞
	//如果客户端关闭了conn测，就不会阻塞
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		return
	}
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkgLen读取消息内容
	n, errs := this.Conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || errs != nil {
		return
	}
	//把pkgLen发序列化成 -->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.r=Write err=", err)
		return
	}
	//发送data 本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.r=dataWrite err=", err)
		return
	}
	return
}
