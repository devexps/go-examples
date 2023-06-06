package bootstrap

import (
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *conf.Data_Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           int(cfg.Db),
		DialTimeout:  cfg.DialTimeout.AsDuration(),
		ReadTimeout:  cfg.ReadTimeout.AsDuration(),
		WriteTimeout: cfg.WriteTimeout.AsDuration(),
	})
	if rdb == nil {
		log.Fatalf("failed opening connection to redis")
	}
	rdb.AddHook(redisotel.NewTracingHook())
	return rdb
}
