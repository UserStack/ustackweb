package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "ustackweb/routers"
	"ustackweb/utils"
)

func main() {
	utils.LoadLocales()
	beego.EnableAdmin = true
	beego.EnableXSRF = true
	beego.XSRFKEY = "63oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.XSRFExpire = 3600
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
