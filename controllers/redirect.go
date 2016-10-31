package controllers

import (
	"fmt"
	"short-url/models"
	"short-url/utils"
	"strings"
)

type RedirectController struct {
	BaseController
}

func (this *RedirectController) Get() {
	if this.Ctx.Request.URL.Path == "/" {
		fmt.Println("RediretController has been used, the url is ", this.Ctx.Request.URL.Path)
		this.Data["title_name"] = "short-link service_beego"
		this.TplName = "index.html"
		return

	}
	urlPath := this.Ctx.Request.URL.Path
	urlPath = strings.Trim(urlPath, "/")
	id := utils.StringToId(urlPath)
	u := &models.Url{
		Id: id,
	}
	err := u.FindById()
	if err != nil {
		this.error("no corresponding short URL in MongoDB:" + err.Error())
		return
	}
	this.Redirect(u.SourceUrl, 302)
}
