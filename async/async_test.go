package async

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//TestGoAsyncRequest go test -v
func TestGoAsyncRequest(t *testing.T) {
	Convey("func GoAsyncRequest()", t, func() {
		Convey("correct", func() {
			res, err := request()

			So(err, ShouldBeNil)
			So(res["a"][0], ShouldEqual, 1)
			So(res["b"][0], ShouldEqual, 2)
			So(res["c"][0], ShouldEqual, 3)
		})
	})
}

//Benchmark_GoAsyncRequest go test -v -bench=".*"
func BenchmarkGoAsyncRequest(b *testing.B) {
	Convey("bench GoAsyncRequest()", b, func() {
		for i := 0; i < b.N; i++ {
			request()
		}
	})
}

func request() (map[string][]interface{}, error) {
	var addFunc MultiAddFunc
	addFunc = append(
		addFunc,
		AddFunc{
			Logo:    "a",
			Handler: requestA,
			Params: []interface{}{
				"str",
			},
		},
	)
	addFunc = append(
		addFunc, AddFunc{
			Logo:    "b",
			Handler: requestB,
		},
	)
	addFunc = append(
		addFunc,
		AddFunc{
			Logo:    "c",
			Handler: requestC,
			Params: []interface{}{
				3,
			},
		},
	)

	return GoAsyncRequest(addFunc, 3)
}

func requestA(str string) int {
	return 1
}

func requestB() int {
	return 2
}

func requestC(i int) int {
	return i
}
