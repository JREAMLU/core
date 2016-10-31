package redigos

import (
	"errors"
	"fmt"
	"time"

	"github.com/JREAMLU/core/io"
	"github.com/garyburd/redigo/redis"

	yaml "gopkg.in/yaml.v2"
)

var redisConf map[string]*RedisConfig

const (
	RedisConfigNotExists = `redis config not exists,server=%s,master=%v`
)

type RedisLoads struct {
	RedisConf []*RedisConfig `yaml:"redisConfig"`
}

type RedisConfig struct {
	Name     string    `yaml:"name"`
	PoolSize int       `yaml:"poolSize"`
	TimeOut  int       `yaml:"timeOut"`
	Connects []Connect `yaml:"connect"`
}

type Connect struct {
	IP     string `yaml:"ip"`
	DB     string `yaml:"db"`
	Master bool   `yaml:"master"`
}

func LoadRedisConfig(filePath string) error {
	redisConf = make(map[string]*RedisConfig)
	var settings RedisLoads
	result, err := io.ReadAllBytes(filePath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(result, &settings); err != nil {
		return err
	}

	if len(settings.RedisConf) == 0 {
		return errors.New("redis load error")
	}

	for i, v := range settings.RedisConf {
		redisConf[v.Name] = settings.RedisConf[i]
	}

	return nil
}

func GetRedisConn(server string, isMaster bool) *Connect {
	if conf, ok := redisConf[server]; ok {
		for i := range conf.Connects {
			if conf.Connects[i].Master == isMaster {
				return &conf.Connects[i]
			}
		}
	}
	return nil
}

func GetRedisClient(server string, isMaster bool) redis.Conn {
	conn := GetRedisConn(server, isMaster)
	if conn == nil {
		return nil
	}
	return GetRedisPool(
		conn.IP,
		conn.DB,
		redisConf[server].PoolSize,
		time.Duration(redisConf[server].TimeOut)*time.Second,
	).Get()
}

func redisConfNotExists(server string, isMaster bool) error {
	return fmt.Errorf(RedisConfigNotExists, server, isMaster)
}
