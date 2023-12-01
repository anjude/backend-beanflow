package redis_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anjude/backend-beanflow/infrastructure/config"

	redis "github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedisClient(config config.RedisConfig) error {
	if !config.Enable {
		log.Printf("redis client is disabled")
		return nil
	}
	if redisClient != nil {
		return fmt.Errorf("redis client already init")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	redisClient = rdb
	return nil
}

func GetRedis() *redis.Client {
	if redisClient == nil {
		log.Fatalf("redis client is nil")
	}
	return redisClient
}

func Get(ctx context.Context, key string) ([]byte, error) {
	return GetRedis().Get(ctx, key).Bytes()
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return GetRedis().Set(ctx, key, value, expiration).Err()
}
