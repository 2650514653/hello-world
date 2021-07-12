package routers

import (
	"beegoProject2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//动态路由
	beego.Router("/api", &controllers.ApiController{})
}
