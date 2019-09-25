package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LogResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"'` //消息类型
	Data string `json:"data"`  //消息内容
}

type LoginMes struct {
	UserId   int    `json:"userid"'`
	UserPwd  string `json:"userpwd"'`
	UserName string `json:"username"'`
}

type LoginResMes struct {
	Code  int    `json:"code"` //状态码 500 该用户没有注册 200： 登录成功
	Error string `json:"error"`
}

type RegisterMes struct {
}
