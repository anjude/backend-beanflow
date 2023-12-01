package wechat

type LoginResp struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
