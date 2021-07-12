package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// c.Data["title"] = "你好,beego"
	// c.Data["num"] = 12
	// c.TplName = "goods.tpl"
	c.Ctx.WriteString("用户中心")
}
func (c *UserController) AddUser() {
	c.TplName = "user.html"

}

//处理post请求
func (c *UserController) DoAddUser() {
	Id, err := c.GetInt("ID")
	if err != nil {
		c.Ctx.WriteString("传入Id格式无效")
		return
	}

	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetStrings("hobby")
	fmt.Print(hobby)

	c.Ctx.WriteString("id:" + strconv.Itoa(Id) + "用户名：" + username + "密码" + password)
}
func (c *UserController) EditUser() {
	c.TplName = "userEdit.html"

}
func (c *UserController) DoEditUser() {
	c.TplName = "userEdit.html"

}

type User struct {
	UserName string
	Password string
	Hobby    []string
}

//在beego中返回json数据
func (c *UserController) GetUser() {
	u := User{
		UserName: "张三",
		Password: "123456",
		Hobby:    []string{"1", "2"},
	}
	//返回一个json数据
	c.Data["json"] = u
	c.ServeJSON()

}
