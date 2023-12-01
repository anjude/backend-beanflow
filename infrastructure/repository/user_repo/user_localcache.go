package user_repo

import (
	"time"

	"github.com/anjude/backend-beanflow/domain/user/do"
	"github.com/anjude/backend-beanflow/infrastructure/beancache"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/pkg/errors"
)

var KeyNotFound = errors.New("key not found")

type UserLocalCacheRepo struct {
}

func (u *UserLocalCacheRepo) GetUserByOpenid(ctx *beanctx.BizContext, openid string) (*do.User, error) {
	user, ok := beancache.Get(ctx, u.GetUserCacheKey(openid))
	if !ok {
		return nil, KeyNotFound
	}
	return user.(*do.User), nil
}

func (u *UserLocalCacheRepo) SetUserCache(ctx *beanctx.BizContext, user *do.User) error {
	beancache.Set(ctx, u.GetUserCacheKey(user.Openid), user, time.Hour*2)
	return nil
}

func (u *UserLocalCacheRepo) GetUserCacheKey(openid string) string {
	return "user:openid:" + openid
}
