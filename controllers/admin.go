package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"hello/models"
	"strings"
	"time"
)

//  AdminController operations for Admin
type AdminController struct {
	baseController
}

func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		pwd := c.GetString("pwd")

		user := models.Admin{Name: name}
		c.o.Read(&user, "name")
		if user.Pwd == "" {
			c.History("账号不存在", "")
		}

		if user.Pwd != strings.Trim(pwd, " ") {
			c.History("密码错误", "")
		}

		user.LastIp = c.getClientIp()
		user.LoginCount = user.LoginCount + 1

		if _, err := c.o.Update(&user); err != nil {
			c.History("登录异常", "")
		} else {
			c.SetSession("user", user)
			c.History("登录成功", c.controllerName+"/index.html")
		}

		//c.Ctx.WriteString(name +"----"+ pwd)

	} else {
		c.TplName = c.controllerName + "/login.html"
	}
}

func (c *AdminController) Add() {
	admin := models.Admin{}
	admin.Name = "admin"
	admin.Pwd = "123456"
	admin.Email = "admin@hello.lmqde.com"
	admin.LastTime = time.Now()
	admin.Created = time.Now()
	admin.Updated = time.Now()
	adminId, err := c.o.Insert(&admin)
	if err != nil {
		c.Ctx.WriteString("fail" + err.Error())
	} else {
		logs.Debug("##########" + fmt.Sprintf("====%v====\n", adminId) + "###########")
		c.Ctx.WriteString("success!adminId:" + fmt.Sprintf("%v", adminId) + "=====" + string(adminId))
	}
}
