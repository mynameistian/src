package routers

import (
	"github.com/astaxie/beego"
	"myBeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about.html", &controllers.AboutController{})
	beego.Router("/contact.html", &controllers.ContactController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/portfolio.html", &controllers.PortfolioController{})
	beego.Router("/services.html", &controllers.ServicesController{})
	beego.Router("/skills.html", &controllers.SkillsController{})
}
