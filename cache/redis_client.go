package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisClients = make(map[string]*redis.Client)
var redisOpts = make(map[string]*redis.Options)

// 创建redis客户端连接实例，并托管
func NewRedisClient(dbIndex, poolSize int, addr, pwd string, alias string) (*redis.Client, error) {
	if alias == "" {
		alias = "default"
	}

	opt := &redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbIndex,
		PoolSize: poolSize,
	}

	return newRedisClient(alias, opt)
}

// 获取redis客户端连接实例
func GetRedisClient() *redis.Client {
	return GetRedisClientAlias("default")
}

// 根据别名获取redis客户端连接实例
func GetRedisClientAlias(alias string) *redis.Client {
	client := redisClients[alias]
	if client == nil || client.PoolStats().TotalConns == 0 {
		opt := redisOpts[alias]
		if opt == nil {
			panic(fmt.Sprintf("redis client [%s] not init", alias))
		} else {
			// 重连
			client, err := newRedisClient(alias, opt)
			if err != nil {
				panic(err)
				return nil
			} else {
				return client
			}
		}
	} else {
		return client
	}
}

func newRedisClient(alias string, opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)
	// 连接测试
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	// 保存配置，重连用
	redisClients[alias] = client
	// 缓存实例，托管
	redisOpts[alias] = opt
	return client, nil
}
