package service

import (
	"github.com/anjude/backend-beanflow/infrastructure/repository/user_repo"
	"github.com/anjude/backend-beanflow/third_party/wechat"
)

func NewUserService() *UserService {
	return &UserService{
		userRepo: user_repo.NewUserRepo(),
		wxApi:    &wechat.Api{},
	}
}
