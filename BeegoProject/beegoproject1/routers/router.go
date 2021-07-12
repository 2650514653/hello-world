package routers

import (
	"beegoproject1/controllers"

	"github.com/astaxie/beego"
) 

func init() {
	//访问get方法
	beego.Router("/", &controllers.MainController{}) 
	
	//访问Article控制器的get方法
	beego.Router("/article", &controllers.ArticleController{})
	//访问Article控制器的AddArticle方法
	beego.Router("/article/add", &controllers.ArticleController{},"get:AddArticle")
	//访问Article控制器的EditArticle方法
	beego.Router("/article/Edit", &controllers.ArticleController{},"get:EditArticle")
	//user路由
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{},"get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{},"post:DoAddUser")
	beego.Router("/user/getuser", &controllers.UserController{},"post:GetUser")
	
	//访问Goods控制器的get方法
	beego.Router("/goods", &controllers.GoodsController{}) 
	//访问Goods控制器的post方法
	beego.Router("/goods/add", &controllers.GoodsController{},"post:DoAdd") 
	//访问Goods控制器的put方法
	beego.Router("/goods/edit", &controllers.GoodsController{},"put:DoEdit") 
	//访问Goods控制器的delete方法
	beego.Router("/goods/delete", &controllers.GoodsController{},"delete:DoDelete") 
	
	beego.Router("/goods/xml", &controllers.GoodsController{},"post:Xml") 

}
