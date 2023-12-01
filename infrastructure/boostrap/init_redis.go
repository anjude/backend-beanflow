package boostrap

import (
	"github.com/anjude/backend-beanflow/infrastructure/client/redis_client"
	"github.com/anjude/backend-beanflow/infrastructure/global"
)

func InitRedis() error {
	return redis_client.InitRedisClient(global.Conf.RedisConfig)
}
