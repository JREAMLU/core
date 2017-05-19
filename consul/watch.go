package consul

import (
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
)

//WatchKey 监听某个具体key变化
// 监听到变化会调用handle，参数为变化key的相关信息
func WatchKey(consulAddr, key string, handle func(*api.KVPair)) {
	plan, err := watch.Parse(map[string]interface{}{
		"type": "key",
		"key":  key,
	})
	if err != nil {
		panic(err)
	}
	first := true
	plan.Handler = func(idx uint64, raw interface{}) {
		// 启动的时候，watch的key如果存在就会调用Handler
		// 所以要跳过
		if first {
			first = false
			return
		}
		if raw == nil {
			return
		}
		v, ok := raw.(*api.KVPair)
		if ok && v != nil {
			handle(v)
		}
	}
	err = plan.Run(consulAddr)
	if err != nil {
		panic(err)
	}
}

//WatchKeyPrefix 监听某一级目录下所有key变化
// 监听到变化会调用handle，参数为变化key的相关信息
//TODO 需要提供stop?
func WatchKeyPrefix(consulAddr, key string, handle func(api.KVPairs)) {
	if !strings.HasSuffix(key, "/") {
		key += "/"
	}

	plan, err := watch.Parse(map[string]interface{}{
		"type":   "keyprefix",
		"prefix": key,
	})
	if err != nil {
		panic(err)
	}
	first := true
	plan.Handler = func(idx uint64, raw interface{}) {
		if first {
			first = false
			return
		}
		if raw == nil {
			return
		}
		v, ok := raw.(api.KVPairs)
		if ok && len(v) > 0 {
			handle(v)
		}
	}

	err = plan.Run(consulAddr)
	if err != nil {
		panic(err)
	}
}
