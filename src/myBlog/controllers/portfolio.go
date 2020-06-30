package controllers

import (
	"github.com/astaxie/beego"
)

type PortfolioController struct {
	beego.Controller
}

func (c *PortfolioController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "portfolio.html"
}
