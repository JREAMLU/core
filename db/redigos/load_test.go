package redigos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadRedisConfig(t *testing.T) {
	Convey("func LoadRedisConfig()", t, func() {
		Convey("correct", func() {
			err := LoadRedisConfig("./redis.yml")
			So(err, ShouldBeNil)
		})
	})
}
