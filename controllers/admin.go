package controllers

import "hello/models"

//  AdminController operations for Admin
type AdminController struct {
	baseController
}

func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")

		user := models.Users{}
	}
}
