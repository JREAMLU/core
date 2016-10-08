package redigos

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	_default_pagesize    = 500
	_default_maxidle     = 50
	_default_idletimeout = 240 * time.Second
)

type RedisStructure struct {
	KeyPrefixFormat string
	ServerName      string
	readPool        *redis.Pool
	writePool       *redis.Pool
	writeConn       string
	readConn        string
}

func NewRedisStructure(serverName, keyPrefixFormat string) RedisStructure {
	return RedisStructure{
		KeyPrefixFormat: keyPrefixFormat,
		ServerName:      serverName,
	}
}

func (rs *RedisStructure) getPool(server string, isMaster bool) *redis.Pool {
	conn := GetRedisConn(server, isMaster)
	if conn == nil {
		return nil
	}
	return GetRedisPool(conn.IP, conn.DB, _default_maxidle, _default_idletimeout)
}

func (rs *RedisStructure) InitPool(connStr, db string) {
	rs.writePool = GetRedisPool(connStr, db, _default_maxidle, _default_idletimeout)
	rs.readPool = rs.writePool
	rs.writeConn = connStr
	rs.readConn = connStr
}

func (rs *RedisStructure) InitRedisKey(keySuffix string) string {
	if keySuffix == "" {
		return rs.KeyPrefixFormat
	} else {
		return fmt.Sprintf(rs.KeyPrefixFormat, keySuffix)
	}
}

func (rs *RedisStructure) getConn(isMaster bool) redis.Conn {
	if rs.writePool == nil {
		rs.writePool = rs.getPool(rs.ServerName, true)
		rs.readPool = rs.getPool(rs.ServerName, false)
		if rs.readPool == nil {
			rs.readPool = rs.writePool
		}
	}

	if isMaster {
		if rs.writePool == nil {
			return nil
		}
		return rs.writePool.Get()
	} else {
		if rs.readPool == nil {
			return nil
		}
		return rs.readPool.Get()
	}
}

func (rs *RedisStructure) String(isMaster bool, cmd string, params ...interface{}) (reply string, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return "", redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.String(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Strings(isMaster bool, cmd string, params ...interface{}) (reply []string, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Strings(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) StringMap(isMaster bool, cmd string, params ...interface{}) (reply map[string]string, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.StringMap(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Int(isMaster bool, cmd string, params ...interface{}) (reply int, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return 0, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Int(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Ints(isMaster bool, cmd string, params ...interface{}) (reply []int, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Ints(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) IntMap(isMaster bool, cmd string, params ...interface{}) (reply map[string]int, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.IntMap(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Int64(isMaster bool, cmd string, params ...interface{}) (reply int64, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return 0, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Int64(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Int64Map(isMaster bool, cmd string, params ...interface{}) (reply map[string]int64, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Int64Map(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Bool(isMaster bool, cmd string, params ...interface{}) (reply bool, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return false, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Bool(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Bytes(isMaster bool, cmd string, params ...interface{}) (reply []byte, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Bytes(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) ByteSlices(isMaster bool, cmd string, params ...interface{}) (reply [][]byte, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.ByteSlices(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Float64(isMaster bool, cmd string, params ...interface{}) (reply float64, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return 0, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Float64(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) MultiBulk(isMaster bool, cmd string, params ...interface{}) (reply []interface{}, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.MultiBulk(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Uint64(isMaster bool, cmd string, params ...interface{}) (reply uint64, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return reply, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Uint64(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Values(isMaster bool, cmd string, params ...interface{}) (reply []interface{}, err error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	reply, err = redis.Values(conn.Do(cmd, params...))
	conn.Close()
	return
}

func (rs *RedisStructure) Keys(pattern string) ([]string, error) {
	return rs.Strings(false, "KEYS", pattern)
}

func (rs *RedisStructure) DelKey(keySuffix string) (bool, error) {
	key := rs.InitRedisKey(keySuffix)
	n, err := rs.Int(true, "DEL", key)
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (rs *RedisStructure) getConnstr(isMaster bool) string {
	if isMaster && rs.writeConn != "" {
		return rs.writeConn
	}
	if !isMaster && rs.readConn != "" {
		return rs.readConn
	}
	redisConn := GetRedisConn(rs.ServerName, isMaster)
	if redisConn == nil {
		redisConn = GetRedisConn(rs.ServerName, !isMaster)
	}
	if redisConn == nil {
		return ""
	}
	if isMaster {
		rs.writeConn = redisConn.IP
	} else {
		rs.readConn = redisConn.IP
	}
	return redisConn.IP
}

func (rs *RedisStructure) Ping(isMaster bool) (interface{}, error) {
	conn := rs.getConn(isMaster)
	if conn == nil {
		return nil, redisConfNotExists(rs.ServerName, isMaster)
	}
	defer conn.Close()
	return conn.Do("PING")
}

func (rs *RedisStructure) Exists(key string) (bool, error) {
	isMaster := false
	conn := rs.getConn(isMaster)
	if conn == nil {
		return false, redisConfNotExists(rs.ServerName, isMaster)
	}
	defer conn.Close()
	exists, err := rs.Int(isMaster, "EXISTS", key)
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (rs *RedisStructure) Rename(keySuffix, newKey string) (string, error) {
	key := rs.InitRedisKey(keySuffix)
	return rs.String(true, "RENAME", key, newKey)
}
