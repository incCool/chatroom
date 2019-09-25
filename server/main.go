package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*4)
	_, err = conn.Read(buf[:4])

	//conn.Read 在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn , 就不会阻塞
	if err != nil {
		//err = errors.New("read pkg header errror!")
		return
	}
	//讲切片转换成 uint32类型 到底要读多少字节
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen 读取内容，从conn里面读pkgLen内容放在buf中
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

//处理登录的请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.将传入的数据反序列化
	var loginMes message.LoginMes

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//声明一个返回结构体
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		fmt.Println("登录成功")
		loginResMes.Code = 200
		loginResMes.Error = ""
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在！"
	}

	//讲loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	//将data赋值给resMes.Data
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail err = ", err)
	}
	//发送给客户端
	writePkg(conn, data)
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var byteIn [4]byte
	binary.BigEndian.PutUint32(byteIn[0:4], pkgLen)

	_, err = conn.Write(byteIn[0:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送Data
	n, err := conn.Write(data)
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Write fail ", err)
		return
	}

	fmt.Println("server response sucess!")

	return
}

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//login
		err = serverProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("serverProcsssLogin err=", err)
			return
		}
	case message.RegisterMesType:
		fmt.Println("注册开始...")

	default:
		fmt.Println("消息类型学不存在....")
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	//read mes from client
	for {
		var mes message.Message
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit! --> server exit!")
				return
			} else {
				fmt.Println("readPkg err=", err)
				return
			}
		}

		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
		}
		fmt.Println("serverProcessMes sucess!")
	}
}

func main() {

	//监听
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err = ", err)
	}
	for {
		fmt.Println("客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept() err=", err)
		}

		//启动协程和客户端保持通讯
		go process(conn)
	}

}
