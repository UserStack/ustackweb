package main

import (
	_ "ustackweb/routers"
	"github.com/astaxie/beego"
  "github.com/beego/i18n"
  "strings"
  "strconv"
)

// Initialized language type list.
func loadLocales() {
  langs := strings.Split("en-US", "|")
  for _, lang := range langs {
      beego.Trace("Loading language: " + lang)
      if err := i18n.SetMessage(lang, "locales/"+lang+".ini"); err != nil {
          beego.Error("Fail to set message file: " + err.Error())
          return
      }
  }
}

func main() {
  loadLocales()
  beego.SessionOn = true
  beego.EnableAdmin = true
  beego.EnableXSRF = true
  beego.XSRFKEY = "63oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
  beego.XSRFExpire = 3600
  beego.AddFuncMap("i18n", i18n.Tr)
  beego.AddFuncMap("intToStr", strconv.Itoa)
	beego.Run()
}

