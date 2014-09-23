package main

import (
	_ "ustackweb/routers"
	"github.com/astaxie/beego"
)

func main() {
  beego.SessionOn = true
  beego.EnableAdmin = true
	beego.Run()
}

