package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {

	//注册数据驱动 mysql sqlite3 postgres 是beego默认注册过的 所以无需设置
	//orm.RegisterDriver("mysql",orm.DRMySQL)

	//注册数据库orm必须注册一个别名为 default的数据库 做默认使用
	//orm.RegisterDataBase("default","mysql","root:123lmqde@tcp(47.52.22.55:3306)/hello")

	//orm.RegisterModel(new(Users))
	//
	//orm.RunSyncdb("default", true, true)

	//1 解析配置
	//2 执行用户hookfunc
	//3 是否开启session
	//4 是否编译模板
	//5 是否开启文档功能
	//6 是否启动管理模块
	//7 监听服务端口
	beego.Run()
}
