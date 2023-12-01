package beanctx

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	Openid *string
	AppId  *string
}

const (
	UserInfoKey = "user_info"
)

type metadata UserInfo

type JwtClaims struct {
	jwt.RegisteredClaims `json:"-"`
	UserInfo
}

func newMetadata(ctx *gin.Context) *metadata {
	m := metadata{}
	if ctx == nil {
		return &m
	}
	value, exists := ctx.Get(UserInfoKey)
	if !exists {
		return &m
	}
	jwtClaims := value.(*JwtClaims)
	m.Openid = jwtClaims.Openid
	m.AppId = jwtClaims.AppId
	return &m
}

func (m *metadata) GetOpenid() string {
	if m.Openid == nil {
		return ""
	}
	return *m.Openid
}

func (m *metadata) GetAppId() string {
	if m.AppId == nil {
		return ""
	}
	return *m.AppId
}
