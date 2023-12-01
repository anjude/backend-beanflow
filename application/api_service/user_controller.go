package api_service

import (
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/domain/user/service"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
)

// 匿名定义，未实现接口可以报错
var _ IUserController = &UserController{}

type IUserController interface {
	Login(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError)
}

type UserController struct {
	userService service.IUserService
}

func (u UserController) Login(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError) {
	req := ctx.GetReqParam().(dto.LoginReq)
	resp, bizErr := u.userService.Login(ctx, req)
	if bizErr != nil {
		return nil, bizErr
	}
	return resp, nil
}
