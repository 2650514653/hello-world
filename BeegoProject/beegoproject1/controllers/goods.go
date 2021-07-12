package controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}
type Product struct {
	Title   string `form:"title xml:"title"`
	Content string `form:"content xml:"content"`
}

func (c *GoodsController) Get() { //get请求
	c.Data["title"] = "你好,beego"
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}
func (c *GoodsController) DoAdd() { //post请求
	c.Ctx.WriteString("执行增加操作")

}
func (c *GoodsController) DoEdit() { //put请求
	c.Ctx.WriteString("执行修改操作")
}
func (c *GoodsController) DoDelete() { //delete
	c.Ctx.WriteString("执行删除操作")

}

//接收post传过来的xml数据
//需要在配置文件里设置copyrequestbody = true
func (c *GoodsController) Xml() { //delete
	var p Product
	str := string(c.Ctx.Input.RequestBody)
	beego.Info(str) //打印获取到的数据

	c.Ctx.WriteString(str)
	var err error
	if e:=xml.Unmarshal(c.Ctx.Input.RequestBody,&p);e!=nil{
		c.Data["json"]=err
		c.ServeJSON()
	}else {
		fmt.Printf("%#v",p)
		c.Data["json"]=p
		c.ServeJSON()
	}

}
