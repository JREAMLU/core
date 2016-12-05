package jlogs

import (
	"time"

	"github.com/astaxie/beego"
)

/*
const (
	LevelTrace         = 7
	LevelDebug         = 7
	LevelInfo          = 6
	LevelInformational = 6
	LevelNotice        = 5
	LevelWarning       = 4
	LevelWarn          = 4
	LevelError         = 3
	LevelCritical      = 2
	LevelAlert         = 1
	LevelEmergency     = 0
)
*/

// InitLogs init logs
func InitLogs() {
	path := beego.AppConfig.String("log.path") + beego.AppConfig.String("appname") + time.Now().Format("2006-01-02") + ".log"
	file := beego.AppConfig.String("log.file")
	console, _ := beego.AppConfig.Bool("log.console")
	level, _ := beego.AppConfig.Int("log.level")
	filename := `{"filename":"` + path + `","separate":["` + file + `"],"daily":false}`

	if !console {
		beego.BeeLogger.DelLogger("console")
	}

	beego.SetLogger("multifile", filename)

	//ElasticSearch
	// esFilename := `{"dsn":"` + beego.AppConfig.String("es.dns") + `"}`
	// fmt.Println(esFilename)
	// beego.SetLogger("es", esFilename)

	beego.SetLevel(level)
}
