package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	// c.Data["title"] = "你好,beego"
	// c.Data["num"] = 12
	// c.TplName = "goods.tpl"
	c.Ctx.WriteString("新闻列表") //直接给页面返回一串字符
}
//给结构体绑定其他方法
func (c *ArticleController) AddArticle() {
	// c.Data["title"] = "你好,beego"
	// c.Data["num"] = 12
	// c.TplName = "goods.tpl"
	c.Ctx.WriteString("添加新闻") //直接给页面返回一串字符
}
func (c *ArticleController) EditArticle() {
	id,err:=c.GetInt("id") //获取传过来的id的值
	if err!=nil {
		beego.Info(err)
	}
	fmt.Printf("传过来的值:%v 类型:%T",id,id)
	c.Ctx.WriteString("修改新闻")
}

