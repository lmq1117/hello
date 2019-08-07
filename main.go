package main

import (
	"github.com/astaxie/beego"
	"hello/models"
	_ "hello/routers"
)

func init() {
	models.Init()
}

func main() {

	//1 解析配置
	//2 执行用户hookfunc
	//3 是否开启session
	//4 是否编译模板
	//5 是否开启文档功能
	//6 是否启动管理模块
	//7 监听服务端口
	beego.Run()
}
