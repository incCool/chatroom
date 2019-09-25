package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	tmpProcessor := &Processor{
		Conn: conn,
	}
	tmpProcessor.Process2()
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
