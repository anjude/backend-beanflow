package api_service

import "github.com/anjude/backend-beanflow/domain/user/service"

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}
