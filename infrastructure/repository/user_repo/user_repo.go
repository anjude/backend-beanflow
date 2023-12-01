package user_repo

import (
	"github.com/anjude/backend-beanflow/domain/user/do"
	"github.com/anjude/backend-beanflow/domain/user/entity"
	"github.com/anjude/backend-beanflow/domain/user/factory"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUserByOpenid(ctx *beanctx.BizContext, openid string) (*do.User, error)
	CreateUser(ctx *beanctx.BizContext, user *do.User) error
}

type ICacheRepo interface {
	GetUserByOpenid(ctx *beanctx.BizContext, openid string) (*do.User, error)
	SetUserCache(ctx *beanctx.BizContext, user *do.User) error
	GetUserCacheKey(openid string) string
}

type UserRepo struct {
	cache ICacheRepo
}

func (u *UserRepo) GetUserByOpenid(ctx *beanctx.BizContext, openid string) (*do.User, error) {
	// 先从缓存中获取
	user, err := u.cache.GetUserByOpenid(ctx, openid)
	if err == nil {
		return user, nil
	}

	// 缓存中不存在，从数据库中获取
	res := &entity.User{}
	err = ctx.GetDb().Where("openid = ?", openid).First(res).Error
	if err != nil {
		// 返回nil，上层处理注册用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// 写入缓存
	user = factory.BuildUserFromEntity(res)
	_ = u.cache.SetUserCache(ctx, user)
	return user, err
}

func (u *UserRepo) CreateUser(ctx *beanctx.BizContext, user *do.User) error {
	res := &entity.User{
		Openid: user.Openid,
	}
	err := ctx.GetDb().Create(res).Error
	if err != nil {
		return err
	}
	user = factory.BuildUserFromEntity(res)

	// 写入缓存
	_ = u.cache.SetUserCache(ctx, user)
	return nil
}
