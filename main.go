package main

import (
	"flag"
	//	"fmt"
	"short-url/controllers"

	"github.com/astaxie/beego"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()

	beego.RESTRouter("api/url", &controllers.UrlController{})

	beego.Router("/*", &controllers.RedirectController{})

	beego.SetStaticPath("/public", "views")

	beego.Run(port)
}
