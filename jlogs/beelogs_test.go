package jlogs

import (
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogs(t *testing.T) {
	Convey("func Logs()", t, func() {
		Convey("correct", func() {
			beego.SetLogger(logs.AdapterConn, `{"net":"udp4","addr":":1200"}`)
			beego.Info("info 123")
		})
	})
}
