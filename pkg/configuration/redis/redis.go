package redis

import (
	"context"
	rc "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"stock/pkg/configuration"
	"time"
)

type Redis struct {
	config     configuration.ServicesApp
	ctx			context.Context
	RC 			*rc.Client
}

func New(config configuration.ServicesApp, ctx context.Context) *Redis {
	return &Redis{
		config: config,
		ctx: ctx,
	}
}

func (a *Redis) InitRedis() {
	rdb := rc.NewClient(&rc.Options{
		Addr:        a.config.Config.Redis.Host,
		Password:    a.config.Config.Redis.Password,
		DB:          a.config.Config.Redis.DB,
		DialTimeout: time.Duration(500) * time.Millisecond,
		ReadTimeout: time.Duration(500) * time.Millisecond,
	})

	_, err := rdb.Ping(a.ctx).Result()
	if err != nil {
		log.Println("Redis error: %s", err.Error())
	} else {
		a.RC = rdb
		log.Println("Redis connected")
	}

}
