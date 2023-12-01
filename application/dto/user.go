package dto

type LoginReq struct {
	Code  string `json:"code"`
	AppId string `json:"app_id"`
}

type LoginResp struct {
	Openid string `json:"openid"`
	Token  string `json:"token"`
}
