package main

import (
	"context"
	"git.gnlab.com/duohao/share.git/hashids"
	redisHelper "git.gnlab.com/duohao/share.git/redis_helper"
	"git.gnlab.com/duohao/share.git/session"
	"log"
)

func init() {
	InitRedisClient(context.Background())
	hashids.Init()
	redisClient := redis.GetClient()
	// 设置 Session RedisClient
	session.SetRedisClient(redisClient)
}

var (
	redis *redisHelper.Client
)

func InitRedisClient(ctx context.Context) {
	var err error
	//初始化redis
	redis, err = redisHelper.NewRedisClient(ctx, "1.117.8.225", 6379, "86a1b907d54bf7010394bf316e183e67", 1, 10)
	if err != nil {
		log.Fatal(err)
	}
}
