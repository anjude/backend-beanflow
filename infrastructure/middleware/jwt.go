package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/beanlog"
	"github.com/anjude/backend-beanflow/infrastructure/global"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

const (
	TokenExpireDuration = time.Hour * 24 * 30

	NoJwtKey = "no_jwt"
)

// GenToken 生成JWT
func GenToken(appId, openid string) (string, error) {
	c := beanctx.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "coder_bean",                                            // 签发人
		},
		UserInfo: beanctx.UserInfo{
			AppId:  &appId,
			Openid: &openid,
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(global.Conf.GetJwtKey())
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*beanctx.JwtClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &beanctx.JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.Conf.GetJwtKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*beanctx.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		if err := ValidJwt(ctx); err != nil {
			resp := genHttpResp(beanerr.JwtTokenError.AppendMsg(err.Error()), nil)
			beanlog.Infof(GetRequestLog(ctx, nil, resp))
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Next() // 后续的处理函数可以用过c.Get("user_info")来获取当前请求的用户信息
	}
}

func ValidJwt(ctx *gin.Context) error {
	// 免校验
	if ctx.GetBool(NoJwtKey) {
		return nil
	}
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return fmt.Errorf("authorization header is empty")
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return fmt.Errorf("authorization format is invalid")
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	jwtClaims, err := ParseToken(parts[1])
	if err != nil {
		return fmt.Errorf("parse token error: %v", err)
	}
	ctx.Set(beanctx.UserInfoKey, jwtClaims)
	return nil
}

// NoJWTAuth 免除JWT的认证
func NoJWTAuth() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(NoJwtKey, true)
		ctx.Next() // 后续的处理函数可以用过c.Get("no_jwt")来获取当前请求的用户信息
	}
}
