package sysProess

import (
	"chatroom.com/message"
	"chatroom.com/server/model"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//登录函数
func LoginFun(data []byte, conn net.Conn) (Outdata []byte, err error) {
	var messageOb message.LoginMes

	err = json.Unmarshal(data, &messageOb)
	if err != nil {
		fmt.Println("json.Unmarshal err is ", err)
		return
	}
	if messageOb.UserId == 0 ||
		messageOb.UserPwd == "" {
		err = errors.New("关键参数为nil")
		return
	}

	userInfoOb := model.NewUserInfo(messageOb.UserId, messageOb.UserPwd, "", "", "", 0)
	if err != nil {
		fmt.Println("model.NewUserInfo err is ", err)
		return
	}

	err = userInfoOb.FindUserInfo()
	if err != nil {
		fmt.Println("userInfoOb.FindUserInfo err is  ", err)
	}

	//初始化用户信息
	OnlineUserInfoOb := NewOnlineUser(messageOb.UserId, messageOb.UserName, 1, conn)
	if err != nil {
		fmt.Println("NewOnlineUser err is ", err)
		return
	}
	//更新用户信息
	err = OnlineList.AddUserInfo(OnlineUserInfoOb)
	if err != nil {
		fmt.Println("AddUserInfo err is ", err)
		return
	}

	Outdata, err = json.Marshal(userInfoOb)
	if err != nil {
		fmt.Println("json.Marshal err is ", err)
		return
	}
	return
}

//注册函数
func RegisterFun(data []byte) (Outdata []byte, err error) {

	var messageOb message.RegisterMes

	err = json.Unmarshal(data, &messageOb)
	if err != nil {
		fmt.Println("json.Unmarshal err is ", err)
		return
	}

	if messageOb.UserId == 0 ||
		messageOb.UserPwd == "" ||
		messageOb.UserName == "" {
		err = errors.New("关键参数为nil")
		return
	}

	userInfoOb := model.NewUserInfo(messageOb.UserId, messageOb.UserPwd, messageOb.UserName,
		messageOb.Sex, messageOb.Addr, messageOb.Age)
	if err != nil {
		fmt.Println("model.NewUserInfo err is ", err)
		return
	}

	err = userInfoOb.AddUserInfo()
	if err != nil {
		fmt.Println("userInfoOb.FindUserInfo err is  ", err)
		return
	}

	//初始化用户信息
	OnlineUserInfoOb := NewOnlineUser(messageOb.UserId, messageOb.UserName, 0, nil)
	if err != nil {
		fmt.Println("NewOnlineUser err is ", err)
		return
	}
	//添加到用户列表中
	err = OnlineList.AddUserInfo(OnlineUserInfoOb)
	if err != nil {
		fmt.Println("AddUserInfo err is ", err)
		return
	}

	Outdata, err = json.Marshal(userInfoOb)
	if err != nil {
		fmt.Println("json.Marshal err is ", err)
		return
	}
	return
}

//返回用户列表
func FriendsList(data []byte) (Outdata []byte, err error) {

	var messageOb string
	var num int

	err = json.Unmarshal(data, &messageOb)
	if err != nil {
		fmt.Println("json.Unmarshal err is ", err)
		return
	}

	var onlinerUserList message.OnlinerUserList

	var tmpOnlineUserInfo message.OnlineUserInfo

	OnlineList := OnlineList.FriendsList()
	num = len(OnlineList) + 1

	onlinerUserList.UserInfo = make(map[int]*message.OnlineUserInfo, num)

	//遍历在线用户列表
	for k, v := range OnlineList {
		tmpOnlineUserInfo.UserId = v.UserId
		tmpOnlineUserInfo.UserName = v.UserName
		tmpOnlineUserInfo.Status = v.Status
		onlinerUserList.UserInfo[k] = &tmpOnlineUserInfo
	}

	Outdata, err = json.Marshal(onlinerUserList)
	if err != nil {
		fmt.Println("json.Marshal err is ", err)
		return
	}
	return
}
