package utils

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

// Initialized language type list.
func LoadLocales() {
	langs := strings.Split("en-US", "|")
	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "locales/"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
