package controllers

import (
	"github.com/astaxie/beego"
)

type SkillsController struct {
	beego.Controller
}

func (c *SkillsController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "skills.html"
}
