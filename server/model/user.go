package model

//定义用户结构体

type User struct {
	UserId   int    `json:"userId"` //用户信息的json 的key与结构体的字段对应的tag名字一致
	UserrPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//
