package user_repo

import (
	"encoding/json"
	"github.com/anjude/backend-beanflow/domain/user/do"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/client/redis_client"
	"time"
)

type UserRedisRepo struct {
}

func (u *UserRedisRepo) GetUserByOpenid(ctx *beanctx.BizContext, openid string) (*do.User, error) {
	userBytes, err := redis_client.Get(ctx, u.GetUserCacheKey(openid))
	if err != nil {
		return nil, err
	}
	user := &do.User{}
	err = json.Unmarshal(userBytes, user)
	return user, err
}

func (u *UserRedisRepo) SetUserCache(ctx *beanctx.BizContext, user *do.User) error {
	return redis_client.Set(ctx, u.GetUserCacheKey(user.Openid), user, time.Hour*2)
}

func (u *UserRedisRepo) GetUserCacheKey(openid string) string {
	return "user:openid:" + openid
}
