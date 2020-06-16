package DBdao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"net/http"
)

//AddSession 向数据库中添加session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values($1,$2,$3);"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)

	if err != nil {
		return err
	}

	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = $1"

	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id ,username,user_id from sessions where session_id =$1;"

	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessID)

	sess := &model.Session{}

	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)

	return sess, nil
}

//IsLogin 判断用户是否已登录
func IsLogin(r *http.Request) (bool, *model.Session) {

	//根据Cookie的Name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookieValue 获取数据库中的session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
