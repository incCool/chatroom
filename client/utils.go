package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*4)
	_, err = conn.Read(buf[:4])

	//conn.Read 在conn没有关闭的情况下，才会产生阻塞
	//如果客户端关闭了conn , 就不会阻塞
	if err != nil {
		//err = errors.New("read pkg header error!")
		return
	}
	var pkgLen uint32
	//切片转换成uint32类型的数， 代表要读取多少长度的数据
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen 读取内容，从conn里面读pkgLen长度的数据内容放在缓存buf中
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
	}

	//把 buf[:pkgLen] 反序列化-> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送长度,把分片数据的长度给pkgLen
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var byteIn [4]byte
	//把pkgLen 使用分片表示，表示写入多少数据
	binary.BigEndian.PutUint32(byteIn[0:4], pkgLen)

	//发送数据的长度
	_, err = conn.Write(byteIn[0:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送数据
	n, err := conn.Write(data)
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Write fail ", err)
		return
	}

	fmt.Println("send data sucess!")

	return
}
