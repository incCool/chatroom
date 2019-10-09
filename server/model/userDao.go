package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个UserDaode 结构体

type UserDao struct {
	pool *redis.Pool
}

//服务器启动时创建一个userDao, 定义成全局变量，直接使用
var (
	MyUserDao *UserDao
)

//使用工厂模式创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//提供什么方法？？？

//根据用户ID 返回User
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//1.去redis查询用户
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	//user = &User{}
	//反序列化
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	return user, nil
}

//完成登录校验
//1.完成对用户的校验
//2.如果用户的id和密码都正确 返回一个user
//3.如果有误，返回一个错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从redis的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//校验密码
	if user.UserrPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return user, err
}
