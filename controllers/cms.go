package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

// 声明了一个控制器MainController
type CMSController struct {
	//内嵌beegoController
	beego.Controller
}

//func init(){
//	orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
//	orm.RegisterModel(new(models.Users))
//	orm.RunSyncdb("default", false, true)
//}

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
	u := new(models.Users)
	u.Name = "刘备"
	u.Age = 66
	u.Pwd = "123456"
	u.Sex = 1
	models.AddUsers(u)
	c.Ctx.WriteString("hello bee ,this is limq @ cms.get")

}
