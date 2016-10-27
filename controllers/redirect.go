package controllers

type RedirectController struct {
	BaseController
}

func (this *RedirectController) Get() {
	this.Redirect("/api/url", 302)

}
