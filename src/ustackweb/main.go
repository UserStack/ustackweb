package main

import (
	_ "ustackweb/routers"
	"github.com/astaxie/beego"
)

func main() {
  beego.SessionOn = true
  beego.EnableAdmin = true
  beego.EnableXSRF = true
  beego.XSRFKEY = "63oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
  beego.XSRFExpire = 3600
	beego.Run()
}

