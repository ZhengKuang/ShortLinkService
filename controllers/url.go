package controllers

import (
	"net/url"
	"short-url/models"
	"short-url/utils"
)

type UrlController struct {
	BaseController
}

func (this *UrlController) Get() {
	sourceUrl := this.Ctx.Input.Query("sourceUrl")
	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		this.error("Url is not a valid url:" + err.Error())
		return
	}
	u := &models.Url{
		SourceUrl: sourceUrl,
	}
	u.GenId()
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()
	this.success(u)
}
