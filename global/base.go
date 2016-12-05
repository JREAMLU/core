package global

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// NestPreparer nest preparer
type NestPreparer interface {
	NestPrepare()
}

// BaseController base controller
type BaseController struct {
	beego.Controller
	i18n.Locale
}

// langType lang type
type langType struct {
	Lang string
	Name string
}

// Lang lang
var Lang string

func init() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang.types"), "|")
	names := strings.Split(beego.AppConfig.String("lang.names"), "|")
	langTypes := make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Info("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "lang/"+"locale_"+lang+".ini"); err != nil {
			beego.Info("Fail to set message file: " + err.Error())
			return
		}
	}
}

// Prepare prepare
func (r *BaseController) Prepare() {
	//Accept-Language
	acceptLanguage := r.Ctx.Request.Header.Get("Accept-Language")
	if len(acceptLanguage) > 4 {
		acceptLanguage = acceptLanguage[:5] // Only compare first 5 letters.
		if i18n.IsExist(acceptLanguage) {
			Lang = acceptLanguage
		}
	}

	if len(Lang) == 0 {
		Lang = "en-US"
	}
}

// Tr translate
func (r *BaseController) Tr(format string) string {
	return i18n.Tr(Lang, format)
}

// GetLang get lang
func GetLang() string {
	return Lang
}
