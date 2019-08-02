package controllers

import (
	"github.com/astaxie/beego"
)

// 声明了一个控制器MainController
type CMSController struct {
	//内嵌beegoController
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("cms", c.Get)
	c.Mapping("h", c.He)
}

// @router /cms [get]
func (c *CMSController) Get() {
	c.Ctx.WriteString("hello bee ,this is limq @ cms.get")
}

// @router /h [get]
func (c *CMSController) He() {
	c.Ctx.WriteString("hello h,这里是向往的生活")
}
