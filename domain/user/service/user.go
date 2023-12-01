package service

import (
	"github.com/anjude/backend-beanflow/application/dto"
	"github.com/anjude/backend-beanflow/domain/user/do"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/middleware"
	"github.com/anjude/backend-beanflow/infrastructure/repository/user_repo"
	"github.com/anjude/backend-beanflow/third_party/wechat"
)

type IUserService interface {
	Login(ctx *beanctx.BizContext, req dto.LoginReq) (interface{}, *beanerr.BizError)
}

type UserService struct {
	userRepo user_repo.IUserRepo
	wxApi    wechat.IWechatApi
}

func (u UserService) Login(ctx *beanctx.BizContext, req dto.LoginReq) (interface{}, *beanerr.BizError) {
	// 从微信获取用户openid
	openid, err := u.wxApi.GetOpenid(ctx, req.AppId, req.Code)
	if err != nil {
		return nil, beanerr.ExternalError.AppendMsg(err.Error())
	}
	// 判断用户是否存在
	userDO, err := u.userRepo.GetUserByOpenid(ctx, openid)
	if err != nil {
		return nil, beanerr.DBError.AppendMsg(err.Error())
	}
	if userDO == nil {
		// 不存在则创建用户
		userDO = &do.User{
			Openid: openid,
		}
		err = u.userRepo.CreateUser(ctx, userDO)
		if err != nil {
			return nil, beanerr.DBError.AppendMsg(err.Error())
		}
	}
	// 生成token
	token, err := middleware.GenToken(req.AppId, userDO.Openid)
	if err != nil {
		ctx.Log().Errorf("gen token error: %v", err)
		return nil, beanerr.InternalError.AppendMsg(err.Error())
	}
	return &dto.LoginResp{
		Token:  token,
		Openid: userDO.Openid,
	}, nil
}
