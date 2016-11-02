package com

import (
	"reflect"
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

func BenchmarkMapMerge(b *testing.B) {
	Convey("bench MapMerge()", b, func() {
		for i := 0; i < b.N; i++ {
			mapmerge()
		}
	})
}

func mapmerge() map[string]interface{} {
	a := make(map[string]interface{})
	b := make(map[string]interface{})
	a["name"] = "jream"
	b["age"] = 18

	return MapMerge(a, b)
}

func TestEqualMapInt(t *testing.T) {
	Convey("func EqualMapInt()", t, func() {
		Convey("correct", func() {
			map1 := map[string]int{"a": 1, "b": 2, "c": 3}
			map2 := map[string]int{"a": 1, "c": 3, "b": 2}
			ok := EqualMapInt(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]int{"a": 11, "b": 2, "c": 3}
			map2 := map[string]int{"a": 1, "c": 3, "b": 2}
			ok := EqualMapInt(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestEqualMapInt64(t *testing.T) {
	Convey("func EqualMapInt64()", t, func() {
		Convey("correct", func() {
			map1 := map[string]int64{"a": 1, "b": 2, "c": 3}
			map2 := map[string]int64{"a": 1, "c": 3, "b": 2}
			ok := EqualMapInt64(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]int64{"a": 11, "b": 2, "c": 3}
			map2 := map[string]int64{"a": 1, "c": 3, "b": 2}
			ok := EqualMapInt64(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestEqualMapString(t *testing.T) {
	Convey("func EqualMapString()", t, func() {
		Convey("correct", func() {
			map1 := map[string]string{"a": "one", "b": "two", "c": "three"}
			map2 := map[string]string{"a": "one", "b": "two", "c": "three"}
			ok := EqualMapString(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]string{"a": "one", "b": "two", "c": "three"}
			map2 := map[string]string{"aa": "one", "b": "two", "c": "three"}
			ok := EqualMapString(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestEqualMapFloat32(t *testing.T) {
	Convey("func EqualMapFloat32()", t, func() {
		Convey("correct", func() {
			map1 := map[string]float32{"a": 1.1, "b": 2.2, "c": 3.3}
			map2 := map[string]float32{"a": 1.1, "c": 3.3, "b": 2.2}
			ok := EqualMapFloat32(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]float32{"a": 1.1, "b": 2.2, "c": 3.3}
			map2 := map[string]float32{"a": 1.11, "c": 3.3, "b": 2.2}
			ok := EqualMapFloat32(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestEqualMapFloat64(t *testing.T) {
	Convey("func EqualMapFloat64()", t, func() {
		Convey("correct", func() {
			map1 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
			map2 := map[string]float64{"a": 1.1, "c": 3.3, "b": 2.2}
			ok := EqualMapFloat64(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
			map2 := map[string]float64{"a": 1.11, "c": 3.3, "b": 2.2}
			ok := EqualMapFloat64(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestEqualMapInterface(t *testing.T) {
	Convey("func EqualMapInterface()", t, func() {
		Convey("correct", func() {
			map1 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}
			map2 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}
			ok := EqualMapInterface(map1, map2)
			So(ok, ShouldBeTrue)
		})

		Convey("uncorrect", func() {
			map1 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}
			map2 := map[string]interface{}{"a": 1.1, "b": "ba", "c": 3}
			ok := EqualMapInterface(map1, map2)
			So(ok, ShouldBeFalse)
		})
	})
}

func BenchmarkEqualMapInt(b *testing.B) {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"a": 1, "c": 3, "b": 2}

	Convey("bench EqualMapInt()", b, func() {
		for i := 0; i < b.N; i++ {
			EqualMapInt(map1, map2)
		}
	})
}

func BenchmarkEqualMapInterface(b *testing.B) {
	map1 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}
	map2 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}

	for i := 0; i < b.N; i++ {
		EqualMapInterface(map1, map2)
	}
}

func BenchmarkEqualMapDeepEqual(b *testing.B) {
	map1 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}
	map2 := map[string]interface{}{"a": 1.1, "b": "b", "c": 3}

	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(map1, map2)
	}
}
