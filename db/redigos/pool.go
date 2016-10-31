package redigos

import (
	"net"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var redisPools map[string]*RedisPool
var poolLock sync.RWMutex

func init() {
	redisPools = make(map[string]*RedisPool)
}

func newRedisPool(addr string, maxIdle int, idleTimeout time.Duration) *RedisPool {
	return &RedisPool{
		addr:        addr,
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		pools:       make(map[string]*redis.Pool),
	}
}

type RedisPool struct {
	pools       map[string]*redis.Pool
	addr        string
	MaxIdle     int
	IdleTimeout time.Duration
	lock        sync.RWMutex
}

func (p *RedisPool) get(db string) (pool *redis.Pool) {
	p.lock.RLock()
	pool = p.pools[db]
	p.lock.RUnlock()
	return
}

// var redisDialCount int64

func (p *RedisPool) register(db string) (pool *redis.Pool) {
	var ok bool
	p.lock.Lock()
	if pool, ok = p.pools[db]; !ok {
		pool = &redis.Pool{
			MaxIdle:     p.MaxIdle,
			MaxActive:   p.MaxIdle,
			IdleTimeout: p.IdleTimeout,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				tcpAddr, err := net.ResolveTCPAddr("tcp", p.addr)
				if err != nil {
					beego.Info("ResolveTCPAddr Error" + err.Error())
					return nil, err
				}
				tc, err := net.DialTCP("tcp", nil, tcpAddr)
				if err != nil {
					beego.Info("DialTCP Error" + err.Error())
					return nil, err
				}
				if err := tc.SetKeepAlive(true); err != nil {
					beego.Info("SetKeepAlive Error" + err.Error())
					return nil, err
				}
				if err := tc.SetKeepAlivePeriod(2 * time.Hour); err != nil {
					beego.Info("SetKeepAlivePeriod Error" + err.Error())
					return nil, err
				}
				c := redis.NewConn(tc, 10*time.Second, 10*time.Second)

				_, err = c.Do("SELECT", db)
				if err != nil {
					beego.Info("cant select" + db + err.Error())
					return nil, err
				}
				beego.Info("RedisDialAgain")
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
		p.pools[db] = pool
	}
	p.lock.Unlock()
	return pool
}

func (p *RedisPool) Get(db string) (pool *redis.Pool) {
	pool = p.get(db)
	if pool != nil {
		return pool
	}
	return p.register(db)
}

func getRedisPool(addr, db string) (pool *redis.Pool) {
	poolLock.RLock()
	if p, ok := redisPools[addr]; ok {
		pool = p.Get(db)
	}
	poolLock.RUnlock()
	return nil
}

func registerRedisPool(addr string, maxIdle int, idleTimeout time.Duration) *RedisPool {
	var (
		p  *RedisPool
		ok bool
	)
	poolLock.Lock()
	if p, ok = redisPools[addr]; !ok {
		p = newRedisPool(addr, maxIdle, idleTimeout)
		redisPools[addr] = p
	}
	poolLock.Unlock()
	return p
}

func GetRedisPool(addr, db string, maxIdle int, idleTimeout time.Duration) *redis.Pool {
	p := getRedisPool(addr, db)
	if p != nil {
		return p
	}
	return registerRedisPool(addr, maxIdle, idleTimeout).Get(db)
}
