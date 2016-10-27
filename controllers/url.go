package controllers

import (
	"github.com/astaxie/beego"
)

type UrlController struct {
	beego.Controller
}

func (c *UrlController) Get() {
	c.Ctx.WriteString("This is UrlController:Get")

}
