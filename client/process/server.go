package process

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

//登录成功后的界面
func ShowMenu() {

	for {
		fmt.Println("------恭喜***登录成功！-----")
		fmt.Println("------1.显示在线好友！-----")
		fmt.Println("------2.发送消息-----")
		fmt.Println("------3.信息列表-----")
		fmt.Println("------4.退出系统-----")
		fmt.Println("请选择（1-4）")
		var key int
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("显示在线用户列表")

		case 2:
			fmt.Println("发送信息")
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入选项不正确，请重新输入...")
		}
	}
}

//处理服务器消息
func serverProcessMes(conn net.Conn) (err error) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("等待客户端向服务器端发送消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("连接失败！ err=", err)
			return err
		}
		fmt.Println("mes = ", mes)
	}

}
