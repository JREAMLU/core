package curl

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRollingCurl(t *testing.T) {
	Convey("func RollingCurl()", t, func() {
		Convey("correct", func() {
			res, err := request()
			So(err, ShouldBeNil)
			fmt.Println("<<<", res)
			fmt.Println(">>>", err)
		})
	})
}

func request() (string, error) {
	res, err := RollingCurl(
		Requests{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/servera.php",
			Header: map[string]string{
				"Content-Type": "application/json;charset=UTF-8;",
			},
			Raw: `{"name":"KII","age":24}`,
		},
	)
	if err != nil {
		return "", err
	}
	return res, nil
}
