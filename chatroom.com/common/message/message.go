package message

//const 消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesTYpe = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

//Message 消息结构体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//LoginMes 登录结构体
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//LoginResMes 登录结构体
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
