package main

import (
	"chatroom/common/message"
	process2 "chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) Process2() (err error) {
	//read mes from client
	for {
		var mes message.Message
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit! --> server exit!")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
		}
		fmt.Println("serverProcessMes sucess!")
	}
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//login
		tmpProc := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = tmpProc.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println("serverProcsssLogin err=", err)
			return
		}
	case message.RegisterMesType:
		fmt.Println("注册开始...")

	default:
		fmt.Println("消息类型不存在....")
	}
	return
}
