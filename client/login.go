package main

import (
	"chatroom/common/message"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//写一个函数，完成后台登录
func login(userId int, userPwd string) (err error) {

	//1.connect server
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err=", err)
	}

	//2. message struct
	var mes message.Message
	mes.Type = message.LoginMesType

	//3. create message
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.loginMes MARSHAL
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. data ->mes.Data
	mes.Data = string(data)

	//6. mes MARSHAL
	marData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	err = writePkg(conn, marData)
	if err != nil {
		fmt.Println("Write err=", err)
		return
	}

	var loginResMes message.LoginResMes
	mess, err := readPkg(conn)
	if err != nil {
		fmt.Println("recv data err = ", err)
		return
	}
	err = json.Unmarshal([]byte(mess.Data), &loginResMes)
	if err != nil {
		fmt.Println("recv data err = ", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("Login Sucess!")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
		err = errors.New("LoginResMes.Code : 500 , error")
	}
	//fmt.Println("RECV DATA:", loginResMes)
	return
}
