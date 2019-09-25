package main

import (
	"fmt"
	"os"
)

/**
登录界面
*/

var userId int
var userPwd string

func main() {

	//定义变量控制循环打印
	var loop bool = true
	var key int

	// 1.登录界面打印
	for loop {
		fmt.Println("---------------欢迎登录聊天系统-------------")
		fmt.Println("---------------1.登录聊天室-----------------")
		fmt.Println("---------------2.注册账号-------------------")
		fmt.Println("---------------3.退出聊天系统---------------")
		fmt.Println("$ 请输入你的选择")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("$$$ 欢迎登录聊天系统！ $$$")
			loop = false
		case 2:
			fmt.Println(" 请注册你的账号 ")
			loop = false
		case 3:
			fmt.Println("$$$ 退出聊天室 $$$")
			os.Exit(0)
		default:
			fmt.Println(" 输入有误请重新输入 ")
		}
	}

	if key == 1 {
		fmt.Println("请输入用户的ID：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码：")
		fmt.Scanf("%s\n", &userPwd)

		login(userId, userPwd)

	} else if key == 2 {
		fmt.Println("-------2.用户注册------")
	}

}
