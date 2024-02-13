package redis_client 

import (
	"api.com/url-short/settings"
	"github.com/redis/go-redis/v9"
	"context"
)

var ctx = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
    Addr:	  settings.GetEnvVar("REDIS_URL"),
    Password: "", 
    DB:		  0, 
})

func SetItem(key string, val string) error {
	return RedisClient.Set(ctx, key, val, 0).Err()
}

func GetItem(key string) string {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		panic(err.Error())
	}
	return val
}