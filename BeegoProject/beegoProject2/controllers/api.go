package controllers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {

	//md5加密
	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("%x", h.Sum(nil))

}
