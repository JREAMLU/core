package redis

import (
	"fmt"
	"testing"
	"time"

	redis "gopkg.in/redis.v4"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	feed *redis.Client
)

var (
	addr     = "172.16.9.221:6391"
	password = ""
	db       = 0
)

func init() {
	feed, _ = InitRedis(addr, password, db)
}

func TestConRedis(t *testing.T) {
	Convey("func ConRedis()", t, func() {
		feed, err := InitRedis(addr, password, db)
		So(err, ShouldBeNil)
		So(feed, ShouldNotBeNil)
		// feed.Close()
	})
}

func TestString(t *testing.T) {
	uid := "1"
	appID := "1"
	supplierID := "1"
	val := "abc"
	key := fmt.Sprintf("push:user:device:%s:%s:%s", uid, appID, supplierID)

	Convey("func StringSet()", t, func() {
		res := feed.Set(key, val, 8000*time.Second)
		So(res.Val(), ShouldEqual, "OK")
	})

	Convey("func StringGet()", t, func() {
		res := feed.Get(key)
		So(res.Val(), ShouldEqual, val)
	})

	Convey("func StringDel()", t, func() {
		res := feed.Del(key)
		So(res.Val(), ShouldEqual, 1)
	})
}

func BenchmarkString(b *testing.B) {
	Convey("bench StringSet()", b, func() {
		for i := 0; i < b.N; i++ {
			uid := "1"
			appID := "1"
			supplierID := "1"
			key := fmt.Sprintf("push:user:device:%s:%s:%s", uid, appID, supplierID)
			feed.Set(key, i, 8000*time.Second)
		}
	})

	Convey("bench StringGet()", b, func() {
		for i := 0; i < b.N; i++ {
			uid := "1"
			appID := "1"
			supplierID := "1"
			key := fmt.Sprintf("push:user:device:%s:%s:%s", uid, appID, supplierID)
			feed.Get(key)
		}
	})
}

func TestHash(t *testing.T) {
	roomid := "100"
	uid := "10900"
	key := "room:user"
	val := map[string]string{roomid: uid}

	Convey("func HashHMSet()", t, func() {
		res := feed.HMSet(key, val)
		So(res.Val(), ShouldEqual, "OK")
	})

	Convey("func HashHMGet()", t, func() {
		res := feed.HMGet(key, roomid)
		So(res.Val()[0], ShouldEqual, uid)
	})

	Convey("func HashHMGet()", t, func() {
		res := feed.HDel(key, roomid)
		So(res.Val(), ShouldEqual, 1)
	})
}

func TestSet(t *testing.T) {
	key := "push:disable:all"
	val := "aaaaa"

	Convey("func SetSAdd()", t, func() {
		res := feed.SAdd(key, val)
		So(res.Val(), ShouldEqual, 1)
	})

	Convey("func SetSMembers()", t, func() {
		res := feed.SMembers(key)
		So(res.Val()[0], ShouldEqual, val)
	})

	Convey("func SetSDiff()", t, func() {
		key2 := "push:disable:all:o"
		res := feed.SDiff(key, key2)
		So(res.Val()[0], ShouldEqual, val)
	})

	Convey("func SetSRem()", t, func() {
		res := feed.SRem(key, val)
		So(res.Val(), ShouldEqual, 1)
	})
}

func TestZset(t *testing.T) {
	key := "push:zset"
	val := redis.Z{
		Score:  1,
		Member: "redis",
	}

	Convey("func ZsetZAdd()", t, func() {
		res := feed.ZAdd(key, val)
		So(res.Val(), ShouldEqual, 1)
	})

	Convey("func ZsetZRange()", t, func() {
		res := feed.ZRange(key, 0, 20)
		So(res.Val()[0], ShouldEqual, val.Member)
	})

	Convey("func ZsetZAdd()", t, func() {
		res := feed.ZRem(key, val.Member.(string))
		So(res.Val(), ShouldEqual, 1)
	})

}
