package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/UserStack/ustackweb/backend"
	_ "github.com/UserStack/ustackweb/routers"
	"github.com/UserStack/ustackweb/utils"
)

func main() {
	backend.Type = backend.Remote
	utils.LoadLocales()
	beego.EnableAdmin = true
	beego.EnableXSRF = true
	beego.XSRFKEY = "63oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.XSRFExpire = 3600
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
