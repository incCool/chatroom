package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	tmpProcessor := &Processor{
		Conn: conn,
	}
	tmpProcessor.Process2()
}

//初始化userDao

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//初始化redis 连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
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
