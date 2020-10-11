package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

//全局变量
var DB *sql.DB
var Config = beego.AppConfig //配置文件对象

/**
*链接数据库
*/
func Connect() {
	//1.读取conf的配置信息
	dbDriver := Config.String("dirverName")
	dbUser := Config.String("db_user")
	dbPassWord := Config.String("db_password")
	dbIP := Config.String("db_ip")
	dbPort := Config.String("db_port")
	dbName := Config.String("db_name")
	//2.组织链接数据库的字符串
	connUrl := dbUser + ":" + dbPassWord + "@tcp(" + dbIP + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	//3.链接数据库
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置文件")
	}
	fmt.Println("数据库连接成功")
	//4.得到数据库连接db并设为全局变量
	DB = db
}
