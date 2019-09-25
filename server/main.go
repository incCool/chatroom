package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn)( mes message.Message, err error){
	buf := make([]byte,1024*4)
	_,err=conn.Read(buf[:4])

	//conn.Read 在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn , 就不会阻塞
	if err!=nil{
		//err = errors.New("read pkg header errror!")
		return
	}
	//讲切片转换成 uint32类型 到底要读多少字节
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
    //根据pkgLen 读取内容，从conn里面读pkgLen内容放在buf中
	n,err:=conn.Read(buf[:pkgLen])

	if uint32(n)!=pkgLen || err!=nil{
		fmt.Println("conn.Read fail err=",err)
	}

	//把 buf[:pkgLen] 反序列化-> message.Message
	err=json.Unmarshal(buf[:pkgLen],&mes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

func process(conn net.Conn){
	defer conn.Close()
	//read mes from client
    for{
    	var mes message.Message
		mes,err:=readPkg(conn)
		if err!=nil{
			if err == io.EOF{
				fmt.Println("client exit! --> server exit!")
				return
			}else {
				fmt.Println("readPkg err=",err)
				return
			}
		}
		fmt.Println("mes=",mes)
	}
}

func main(){

	//监听
	fmt.Println("服务器在8889端口监听....")
	listen,err:=net.Listen("tcp","127.0.0.1:8889")
    defer listen.Close()
	if err!=nil{
		fmt.Println("net.listen err = ",err)
	}
	for{
		fmt.Println("客户端来连接...")
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("Listen.Accept() err=",err)
		}

		//启动协程和客户端保持通讯
		go process(conn)
	}

}
