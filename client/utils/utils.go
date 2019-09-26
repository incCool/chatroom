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
	Buf  [1024 * 4]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 1024*4)
	_, err = this.Conn.Read(this.Buf[:4])

	//conn.Read 在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn , 就不会阻塞
	if err != nil {
		//err = errors.New("read pkg header errror!")
		return
	}
	//讲切片转换成 uint32类型 到底要读多少字节
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen 读取内容，从conn里面读pkgLen内容放在buf中
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
	}

	//把 buf[:pkgLen] 反序列化-> message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var byteIn [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	_, err = this.Conn.Write(this.Buf[0:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送Data
	n, err := this.Conn.Write(data)
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Write fail ", err)
		return
	}

	//fmt.Println("[ NEW ] server response sucess!")

	return
}
