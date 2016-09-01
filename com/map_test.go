package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMapMerge(t *testing.T) {
	Convey("func MapMerge()", t, func() {
		Convey("correct", func() {
			c := mapmerge()

			So(c, ShouldContainKey, "name")
			So(c, ShouldContainKey, "age")
		})
	})
}

func Benchmark_MapMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapmerge()
	}
}

func mapmerge() map[string]interface{} {
	a := make(map[string]interface{})
	b := make(map[string]interface{})
	a["name"] = "jream"
	b["age"] = 18

	return MapMerge(a, b)
}
