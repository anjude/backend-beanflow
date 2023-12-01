package api_service

import (
	"github.com/anjude/backend-beanflow/domain/flow/flow_service"
	"github.com/anjude/backend-beanflow/domain/user/user_service"
)

func NewUserController() *UserController {
	return &UserController{
		userService: user_service.NewUserService(),
	}
}

func NewFlowController() *FlowController {
	return &FlowController{
		flowService: flow_service.NewFlowService(),
	}
}
