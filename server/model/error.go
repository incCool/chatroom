package model

import "errors"

//定义用户登录时的错误类型

var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在!!!")
	ERROR_USER_EXISTS    = errors.New("用户已存在！！！")
	ERROR_USER_PWD       = errors.New("密码不正确")
)

//用户登录成功

//用户登录失败 -- 密码错误

//用户登录失败 --
