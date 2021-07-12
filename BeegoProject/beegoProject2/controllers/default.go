package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type Article struct {
	Title   string
	Content string
}

func (c *MainController) Get() {
	//在模板中绑定基本的数据类型
	c.Data["title"] = "hello,beego"
	c.Data["num"] = 100
	c.Data["flag"] = true

	//在模板中绑定结构体数据
	article := Article{Title: "golang教程", Content: "beego项目"}
	c.Data["article"] = article
	//在模板中循环遍历切片
	c.Data["sliceList"] = []string{"php", "java", "golang"}

	//在模板中循环遍历map
	userinfo :=make(map[string]interface{},)
	userinfo["username"] ="辉常人"
	userinfo["age"] = 21
	userinfo["sex"] = "男"
	c.Data["userinfo"] = userinfo

	
	c.TplName = "index.tpl"
}
