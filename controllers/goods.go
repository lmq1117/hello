package controllers

import (
	"github.com/astaxie/beego"
)

//  GoodsController operations for Goods
type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
