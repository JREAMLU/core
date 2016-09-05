package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIp2Int(t *testing.T) {
	Convey("func Ip2Int()", t, func() {
		Convey("correct", func() {
			ip := "255.255.255.255"
			So(Ip2Int(ip), ShouldEqual, 4294967295)
		})

		Convey("incorrect", func() {
			ip := "255.255.255.255.255"
			So(Ip2Int(ip), ShouldEqual, 0)
		})
	})
}

func BenchmarkIp2Int(b *testing.B) {
	Convey("bench Ip2Int()", b, func() {
		for i := 0; i < b.N; i++ {
			ip := "255.255.255.255"
			ipInt := Ip2Int(ip)
			Int2Ip(ipInt)
		}
	})
}
