package com

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	err := InitIP2Region("ip2region.db")
	if err != nil {
		log.Println(err)
	}

	Convey("func Query()", t, func() {
		Convey("correct", func() {
			ip, err := Query([]string{"127.0.0.1", "119.75.218.70"}, "memory")
			So(err, ShouldBeNil)
			So(ip, ShouldNotBeNil)
		})
	})
}
