package redis

import (
	"errors"

	"gopkg.in/redis.v4"
)

var R *redis.Client

func InitRedis(addr string, password string, db int) error {
	R = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pong, err := R.Ping().Result()
	if err != nil {
		return err
	}

	if pong == "PONG" {
		return nil
	}

	return errors.New(pong)
}
