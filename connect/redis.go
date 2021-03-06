package connect

import (
	"github.com/astaxie/beego/logs"
	"github.com/brbgithub/goim/conf"
	"github.com/go-redis/redis"
)

const deviceIdPre = "connect:device_id:"

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr: conf.RedisIP,
			DB:0,
		},
	)

	_,err := redisClient.Ping().Result()
	if err != nil {
		logs.Error("redis err ")
		panic(err)
	}
}
