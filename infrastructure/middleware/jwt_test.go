package middleware

import (
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/utils"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenToken(t *testing.T) {
	c := beanctx.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "coder-bean",                                            // 签发人
		},
		UserInfo: beanctx.UserInfo{
			AppId:  utils.GetPtr("your_app_id"),
			Openid: utils.GetPtr("coder_bean"),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	println(token.SignedString([]byte("astest")))
}
