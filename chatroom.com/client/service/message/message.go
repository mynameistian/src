package message

//const 消息类型
const (
	LoginMesTYpe    = "LoginMes"
	LoginResMesTYpe = "LoginResMes"
)

//Message 消息结构体
type Message struct {
	MessageType string `json:"messagetype"`
	Data        string `json:"data"`
}

//LoginMes 登录结构体
type LoginMes struct {
	UserName string `json:"username"`
	UserPwd  string `json:"userpwd"`
}

//LoginResMes 登录结构体
type LoginResMes struct {
	Pause string `json:"pause"`
}
