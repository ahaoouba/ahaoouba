package models

import (
	"fmt"

	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var dbReady bool

// register database orm object
func InitDB() {
	if dbReady {
		return
	}
	beego.Debug("正在初始化数据库连接信息...")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	dburl := "%s:%s@tcp(%s:%d)/%s?charset=%s"
	//"172.17.0.2",
	dburl = fmt.Sprintf(dburl,
		"root",
		"123456",

		"127.0.0.1",
		3306,
		"ahaooubabk",
		"utf8")

	beego.Informational(fmt.Sprintf("数据库连接信息为： %s", dburl))

	err = orm.RegisterDataBase("default", "mysql", dburl)
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	orm.Debug = false

	dbReady = true
}
