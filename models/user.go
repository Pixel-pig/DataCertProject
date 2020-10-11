package models

import (
	"DaraCertProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

/**
 * 存数据
 */
func (u User) SaveUserInfo() (int64, error) {
	//1.将密码进行md5脱敏
	m := md5.New()
	m.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(m.Sum(nil))

	//2.将数据存入数据库中(返回-1则代表错误)
	sqlStr := "insert into user(userPassword, phone) values(?, ?)"
	row, err := db_mysql.DB.Exec(sqlStr, u.Password, u.Phone)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 *查询数据是否存在
 */
func (u User) QueryUserInfo() (error) {
	//1.将密码进行脱敏，（密码存入时进行了加密处理）
	m := md5.New()
	m.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(m.Sum(nil))

	//2.查询数据是否存在
	sqlStr := "select phone from user where phone = ?, userPassword = ?"
	row := db_mysql.DB.QueryRow(sqlStr, u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return err
	}
	return nil
}
