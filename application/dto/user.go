package dto

// LoginReq 为了兼容测试时没有app id，这里不做校验
type LoginReq struct {
	Code  string `form:"code"`
	AppId string `form:"app_id"`
}

type LoginResp struct {
	Openid string `json:"openid"`
	Token  string `json:"token"`
}
