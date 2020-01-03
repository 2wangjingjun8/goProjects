package message

const(
	LoginMesType = "LoginMes" 
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct{
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //消息内容
}

type LoginMes struct{
	UserId int `json:"userId"`	// 用户id
	UserPwd string `json:"userPwd"` // 用户密码
	UserName string `json:"userName"` // 用户名字
}

type LoginResMes struct{
	Code int `json:"code"` // 错误状态码
	Error string `json:"error"` // 错误消息
}

type RegisterMes struct{
	UserId int `json:"userId"`	// 用户id
	UserPwd string `json:"userPwd"` // 用户密码
	UserName string `json:"userName"` // 用户名字
	TelePhone string `json:"telePhone"` // 电话
}