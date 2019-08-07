package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Users), new(Admin))
	orm.Debug = true
	//orm.RunSyncdb("default", true, true)

	//注册数据驱动 mysql sqlite3 postgres 是beego默认注册过的 所以无需设置
	//orm.RegisterDriver("mysql",orm.DRMySQL)

	//注册数据库orm必须注册一个别名为 default的数据库 做默认使用
	//orm.RegisterDataBase("default","mysql","root:123lmqde@tcp(47.52.22.55:3306)/hello")

	//orm.RegisterModel(new(Users))
	//
	//orm.RunSyncdb("default", true, true)
}
