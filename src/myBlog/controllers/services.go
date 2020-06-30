package controllers

import (
	"github.com/astaxie/beego"
)

type ServicesController struct {
	beego.Controller
}

func (c *ServicesController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "services.html"
}
