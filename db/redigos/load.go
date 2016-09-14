package redigos

import (
	"errors"
	"fmt"

	"github.com/JREAMLU/core/io"

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
	Connects []Connect `yaml:"connect"`
}

type Connect struct {
	IP     string `yaml:"ip"`
	Db     int    `yaml:"db"`
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

func redisConfNotExists(server string, isMaster bool) error {
	return fmt.Errorf(RedisConfigNotExists, server, isMaster)
}
