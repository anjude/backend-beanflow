package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/global"
	"github.com/anjude/backend-beanflow/infrastructure/utils/http_util"
	"net/http"
)

var _ IWechatApi = &Api{}

type IWechatApi interface {
	GetOpenid(ctx *beanctx.BizContext, appId, code string) (string, error)
}

type Api struct {
}

func (a *Api) GetOpenid(ctx *beanctx.BizContext, appId, code string) (string, error) {
	// 没有配置appId，使用mock openid
	if global.Conf.App.AppId == "your_app_id" {
		ctx.Log().Warnf("use mock openid")
		return "微信公众号【豆小匠Coding】", nil
	}
	secret := global.Conf.App.AppSecret
	res, err := http_util.DefaultClient.DoReq(http.MethodGet, fmt.Sprintf(getOpenid, appId, secret, code), nil, nil)
	if err != nil {
		return "", err
	}
	resp := LoginResp{}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return "", err
	}
	if resp.Errcode != 0 {
		return "", fmt.Errorf(resp.Errmsg)
	}
	return resp.Openid, nil
}
