package DBdao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {

	sqlStr := "select id ,username,password,email from users where username = $1 and password = $2 ;"

	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {

	sqlStr := "select id ,username,password,email from users where username = $1 ;"

	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil

}

func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username ,password, email)values($1,$2,$3);"

	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
