package factory

import (
	"github.com/anjude/backend-beanflow/domain/user/do"
	"github.com/anjude/backend-beanflow/domain/user/entity"
)

func BuildUserFromEntity(user *entity.User) *do.User {
	return &do.User{
		ID:         user.ID,
		Openid:     user.Openid,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
}
