package main

import (
	"os"

	"github.com/UserStack/ustackweb/backend"
	_ "github.com/UserStack/ustackweb/routers"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
)

func main() {
	backend.Type = backend.Remote
	utils.LoadLocales()
	beego.EnableAdmin = true
	beego.EnableXSRF = true
	beego.XSRFKEY = "63oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.XSRFExpire = 3600
	beego.SessionDomain = os.Getenv("COOKIE_DOMAIN")
	beego.Run()
}
