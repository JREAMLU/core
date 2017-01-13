package jlogs

import (
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

// func TestLogs(t *testing.T) {
// 	Convey("func Logs()", t, func() {
// 		Convey("correct", func() {
// 			filename := `{"filename":"1.log", "separate":[""], "daily": true, "maxdays": 7}`
// 			beego.SetLogger(logs.AdapterMultiFile, filename)
// 			beego.Info("info 123")
// 		})
// 	})
// }

func TestLogs(t *testing.T) {
	Convey("func Logs()", t, func() {
		Convey("correct", func() {
			beego.SetLogger(logs.AdapterConn, `{"net":"udp4","addr":"127.0.0.1:1200"}`)
			beego.Info("info 123")
		})
	})
}

// func BenchmarkLogs(b *testing.B) {
// 	filename := `{"filename":"1.log", "separate":[""], "daily": true, "maxdays": 7}`
// 	beego.SetLogger(logs.AdapterMultiFile, filename)
// 	b.StopTimer()
// 	b.StartTimer()
// 	Convey("bench Logs()", b, func() {
// 		for i := 0; i < b.N; i++ {
// 			beego.Info("info 123")
// 		}
// 	})
// 	b.StopTimer()
// }

// func BenchmarkLogs(b *testing.B) {
// 	beego.SetLogger(logs.AdapterConn, `{"net":"udp4","addr":"127.0.0.1:1200"}`)
// 	b.StopTimer()
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		beego.Info("info 123")
// 	}
// 	b.StopTimer()
// }

// func BenchmarkLogsParallel(b *testing.B) {
// 	beego.SetLogger(logs.AdapterConn, `{"net":"udp4","addr":":1200"}`)
// 	b.StopTimer()
// 	b.StartTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			beego.Info("info 123")
// 		}
// 	})
// 	b.StopTimer()
// }
