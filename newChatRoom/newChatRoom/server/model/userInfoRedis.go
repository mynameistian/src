package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

//用户信息
type UserInfo struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	Sex      string `json:"sex"`
	Addr     string `json:"addr"`
	Age      int    `json:"age"`
}

//工厂模式: 创建用户
func NewUserInfo(userId int, userPwd string, userName,
	sex string, addr string, age int) (userInfoOb *UserInfo) {

	userInfoOb = &UserInfo{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
		Sex:      sex,
		Addr:     addr,
		Age:      age,
	}

	return
}

//添加用户信息
func (userinfo *UserInfo) AddUserInfo() (err error) {

	conn, err := GetPoolConn()
	if err != nil {
		return
	}
	defer conn.Close()

	data, err := json.Marshal(userinfo)
	if err != nil {
		return
	}

	_, err = conn.Do("Hset", userinfo.UserId, data)
	if err != nil {
		return
	}
	return
}

//更新用户信息
func (userinfo *UserInfo) UpdateUserInfo() (err error) {

	if err != userinfo.FindUserInfo() {
		return
	}

	if err != userinfo.AddUserInfo() {
		return
	}

	return
}

//删除用户信息
func (userinfo *UserInfo) DeleteUserInfo() (err error) {

	if err != userinfo.FindUserInfo() {
		return
	}

	conn, err := GetPoolConn()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Do("Hdel", userinfo.UserId)
	if err != nil {
		return
	}

	return
}

//查找用户信息
func (userinfo *UserInfo) FindUserInfo() (err error) {

	conn, err := GetPoolConn()
	if err != nil {
		return
	}
	defer conn.Close()

	Data, err := redis.String(conn.Do("Hget", userinfo.UserId))
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(Data), userinfo)
	return
}
