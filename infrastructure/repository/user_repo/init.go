package user_repo

import (
	"github.com/anjude/backend-beanflow/infrastructure/global"
)

func NewUserRepo() *UserRepo {
	userRepo := &UserRepo{
		cache: &UserRedisRepo{},
	}
	// 如果没有开启redis，则使用本地缓存
	if !global.Conf.RedisConfig.Enable {
		userRepo.cache = &UserLocalCacheRepo{}
	}
	return userRepo
}
