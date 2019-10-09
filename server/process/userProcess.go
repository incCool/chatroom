package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//处理登录的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	//redis数据库用户验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Error = err.Error()
			loginResMes.Code = 500
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Error = err.Error()
			loginResMes.Code = 300
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器错误！！！"
		}
	} else {
		fmt.Println(user, "登录成功")
		loginResMes.Code = 200
		//
	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	fmt.Println("登录成功")
	//	loginResMes.Code = 200
	//	loginResMes.Error = ""
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在！"
	//}

	//讲loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//将data赋值给resMes.Data
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail err = ", err)
		return
	}

	transf := &utils.Transfer{
		Conn: this.Conn,
	}
	//发送给客户端
	transf.WritePkg(data)
	return
}
