package com

import (
	"fmt"
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
			fmt.Println(ip, err)
		})
	})
}

// func BenchmarkMapMerge(b *testing.B) {
// 	Convey("bench Query()", b, func() {
// 		for i := 0; i < b.N; i++ {
// 			mapmerge()
// 		}
// 	})
// }
