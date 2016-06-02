package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const _testKey = "abcdefghijklmnopqrstuvwyxz"

var _testSrc = []byte("123456789")

func TestHMacSha1(t *testing.T) {
	Convey("Test hmac-sha1", t, func() {
		result, err := HMacSha1(_testKey, _testSrc)
		So(err, ShouldBeNil)
		t.Log(result)
	})
}

func TestHMacMd5(t *testing.T) {
	Convey("Test hmac-sha1", t, func() {
		result, err := HMacMD5(_testKey, _testSrc)
		So(err, ShouldBeNil)
		t.Log(result)
	})
}

func TestMd5(t *testing.T) {
	Convey("Test hmac-sha1", t, func() {
		result, err := MD5(string(_testSrc))
		So(err, ShouldBeNil)
		t.Log(result)
	})
}

func TestSha1(t *testing.T) {
	Convey("Test hmac-sha1", t, func() {
		result, err := Sha1(string(_testSrc))
		So(err, ShouldBeNil)
		t.Log(result)
	})
}
