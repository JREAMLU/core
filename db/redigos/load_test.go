package redigos

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/garyburd/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

var serverName = "feed"

func TestLoadRedisConfig(t *testing.T) {
	Convey("func LoadRedisConfig()", t, func() {
		Convey("correct", func() {
			err := LoadRedisConfig("./redis.yml")
			So(err, ShouldBeNil)
		})
	})
}

func TestGetRedisConn(t *testing.T) {
	Convey("func GetRedisConn()", t, func() {
		Convey("correct", func() {
			conn := GetRedisConn(serverName, true)
			So(conn, ShouldNotBeNil)
		})
	})
}

func TestGetRedisClient(t *testing.T) {
	Convey("func GetRedisClient()", t, func() {
		Convey("correct", func() {
			conn := GetRedisClient(serverName, true)
			So(conn, ShouldNotBeNil)
		})
	})
}

func TestInsert(t *testing.T) {
	Convey("func insert()", t, func() {
		Convey("correct", func() {
			reply, err := insert("a")
			So(reply, ShouldBeGreaterThanOrEqualTo, 0)
			So(err, ShouldBeNil)
		})
	})
}

func BenchmarkInsert(b *testing.B) {
	conn := GetRedisClient(serverName, true)
	key := fmt.Sprintf("zgo:%v", "n1")
	defer conn.Close()
	b.StartTimer()
	Convey("bench insert()", b, func() {
		for i := 0; i < b.N; i++ {
			redis.Int64(conn.Do("SADD", key, strconv.Itoa(i)))
		}
	})
	b.StopTimer()
}

func insert(str string) (int64, error) {
	conn := GetRedisClient(serverName, true)
	key := fmt.Sprintf("zgo:%v", "n1")
	reply, err := redis.Int64(conn.Do("SADD", key, str))
	conn.Close()
	return reply, err
}
