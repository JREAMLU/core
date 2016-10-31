package curl

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRollingCurl(t *testing.T) {
	Convey("func RollingCurl()", t, func() {
		Convey("correct", func() {
			res, err := request(15)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeEmpty)
		})

		Convey("uncorrect", func() {
			res, err := request(1)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeEmpty)
		})
	})
}

func request(timeout int64) (string, error) {
	res, err := RollingCurl(
		Requests{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/servera.php",
			Header: map[string]string{
				"Content-Type": "application/json;charset=UTF-8;",
			},
			Raw:        `{"name":"KII","age":24}`,
			RetryTimes: 3,
			Timeout:    timeout,
		},
	)
	if err != nil {
		return "", err
	}
	return res.Body, nil
}
