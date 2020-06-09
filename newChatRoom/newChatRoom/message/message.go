package message

//const 消息类型
const (
	LoginMesType          = "LoginMes"
	LoginResMesTYpe       = "LoginResMes"
	RegisterMesType       = "RegisterMes"
	RegisterinResMesTYpe  = "RegisterinResMes"
	RrivateChatMesType    = "RrivateChatMes"
	GroupChatMesType      = "GroupChatMes"
	RrivateChatResMesType = "RrivateChatResMes"
	GroupChatResMesType   = "GroupChatResMes"
	SysErrResMesType      = "SysErrResMes"
	FriendsListMesType    = "FriendsListMes"
	FriendsListResMesType = "FriendsListResMes"
)

//Message 消息结构体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Code int    `json:"code"`
}

//LoginMes 登录结构体
type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

//RegisterMes 注册结构体
type RegisterMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
	Sex      string `json:"sex"`
	Addr     string `json:"addr"`
	Age      int    `json:"age"`
}

//单个用户信息节点
type OnlineUserInfo struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	Status   int    `jsong:"status"`
}

//在线用户map
type OnlinerUserList struct {
	UserInfo map[int]*OnlineUserInfo `json:"userinfo"`
}

//聊天信息
type MessageData struct {
	SenderId     int    `json:"senderid"`
	SenderName   string `json:"sendername"`
	ReceiverId   int    `jsong:"receiverid"`
	ReceiverName string `jsong:"receivername"`
	Messdata     string `json:"messdata"`
}
