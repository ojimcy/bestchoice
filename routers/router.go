package routers

import (
	"tronbooth/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/books/:address", &controllers.MainController{}, "*:Books")
	beego.Router("/books", &controllers.MainController{}, "*:Books")

	beego.Router("/contact/:address", &controllers.MainController{}, "*:Contact")
	beego.Router("/contact", &controllers.MainController{}, "*:Contact")

	beego.Router("/author/:address", &controllers.MainController{}, "*:Author")
	beego.Router("/author", &controllers.MainController{}, "*:Author")

	beego.Router("/buy/:address", &controllers.MainController{}, "*:Buy")
	beego.Router("/buy", &controllers.MainController{}, "*:Buy")

	beego.Router("/:address", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})

}
