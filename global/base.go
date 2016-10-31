package global

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	i18n.Locale
}

type langType struct {
	Lang string
	Name string
}

var Lang string

func (this *BaseController) Prepare() {
	//Accept-Language
	acceptLanguage := this.Ctx.Request.Header.Get("Accept-Language")
	if len(acceptLanguage) > 4 {
		acceptLanguage = acceptLanguage[:5] // Only compare first 5 letters.
		if i18n.IsExist(acceptLanguage) {
			Lang = acceptLanguage
		}
	}

	if len(Lang) == 0 {
		Lang = "en-US"
	}

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

func GoGetAllAppConfig() map[string]interface{} {
	var core = make(map[string]interface{})
	core["author"] = beego.AppConfig.String("core.author")
	core["ProjectName"] = beego.AppConfig.String("core.ProjectName")

	for key, value := range core {
		fmt.Println(key, ":", value)
	}

	return core
}

func (this *BaseController) GoGetAuthor() string {
	return beego.AppConfig.String("core.author")
}

func (this *BaseController) GoGetProjectName() string {
	return beego.AppConfig.String("core.ProjectName")
}

func (this *BaseController) GogetUrlDomain() string {
	return beego.AppConfig.String("ShortUrl")
}

func (this *BaseController) Tr(format string) string {
	return i18n.Tr(Lang, format)
}

func GetLang() string {
	return Lang
}
