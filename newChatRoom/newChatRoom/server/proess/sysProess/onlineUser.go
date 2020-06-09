package sysProess

import (
	"net"
)

//全局变量
var (
	OnlineList *OnlinerUserList
)

//单个用户信息节点
type OnlineUserInfo struct {
	UserId   int
	UserName string
	Status   int
	Conn     net.Conn
}

//在线用户map
type OnlinerUserList struct {
	UserInfo map[int]*OnlineUserInfo
}

//单例 初始化单个用户信息
func NewOnlineUser(userId int, userName string, status int, conn net.Conn) (OnlineUserOb *OnlineUserInfo) {
	OnlineUserOb = &OnlineUserInfo{
		UserId:   userId,
		UserName: userName,
		Status:   status,
		Conn:     conn,
	}
	return
}

//默认创建map 大小为1024
func InitOnlineList() {
	OnlineList = &OnlinerUserList{
		UserInfo: make(map[int]*OnlineUserInfo, 1024),
	}
}

//添加\更新用户信息到在线列表
func (onlinerUserList *OnlinerUserList) AddUserInfo(onlineUserInfoOb *OnlineUserInfo) (err error) {

	onlinerUserList.UserInfo[onlineUserInfoOb.UserId] = onlineUserInfoOb
	return
}

//删除用户信息
func (onlinerUserList *OnlinerUserList) DeleteUserInfo(onlineUserInfoOb *OnlineUserInfo) (err error) {

	delete(onlinerUserList.UserInfo, onlineUserInfoOb.UserId)
	return
}

//返回整个map
func (onlinerUserList *OnlinerUserList) FriendsList() map[int]*OnlineUserInfo {
	return onlinerUserList.UserInfo
}

//返回整个map
func (onlinerUserList *OnlinerUserList) FindUserInfo(UserId int) *OnlineUserInfo {
	return onlinerUserList.UserInfo[UserId]
}
