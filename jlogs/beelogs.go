package jlogs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// InitLogs init logs
func InitLogs() {
	path := beego.AppConfig.String("log.path") + beego.AppConfig.String("appname") + ".log"
	file := beego.AppConfig.String("log.file")
	console, _ := beego.AppConfig.Bool("log.console")
	level, _ := beego.AppConfig.Int("log.level")
	daily := beego.AppConfig.String("log.daily")
	maxdays := beego.AppConfig.String("log.maxdays")
	filename := `{"filename":"` + path + `", "separate":["` + file + `"], "daily": ` + daily + `, "maxdays": ` + maxdays + `}`

	if !console {
		beego.BeeLogger.DelLogger("console")
	}

	beego.SetLogger(logs.AdapterMultiFile, filename)
	beego.SetLevel(level)
}
