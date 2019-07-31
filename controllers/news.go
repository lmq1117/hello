package controllers

import (
	"github.com/astaxie/beego"
)

// NewsController operations for News
type NewsController struct {
	beego.Controller
}

func (c *NewsController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @route /news [get]
func (c *NewsController) Get() {
	c.Data["Website"] = "beego.news"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
