package controllers

import (
	"github.com/astaxie/beego"
)

// 声明了一个控制器MainController
type MainController struct {
	//内嵌beegoController
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me2"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index2.tpl"
	c.Ctx.WriteString("hello bee ,this is limq")
}
