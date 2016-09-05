package redis

import (
	"errors"

	"gopkg.in/redis.v4"
)

func InitRedis(addr string, password string, db int) (r *redis.Client, err error) {
	r = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pong, err := r.Ping().Result()
	if err != nil {
		return nil, err
	}

	if pong == "PONG" {
		return r, nil
	}

	return r, errors.New(pong)
}
