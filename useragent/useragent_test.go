package useragent

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkUseragent(b *testing.B) {
	var str = `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 QIHU 360EE`
	Convey("bench Useragent()", b, func() {
		for i := 0; i < b.N; i++ {
			ParseByString(str)
		}
	})
}
