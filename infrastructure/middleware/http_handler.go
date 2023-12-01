package middleware

import (
	"fmt"
	"github.com/anjude/backend-beanflow/infrastructure/beanctx"
	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/utils/json_util"
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	ErrCode int32       `json:"err_code"`
	Msg     string      `json:"msg"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
}

func HandleRequest(handler func(ctx *beanctx.BizContext) (interface{}, *beanerr.BizError), req ...interface{}) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var resp *HttpResponse
		var bizErr *beanerr.BizError
		bizCtx := beanctx.NewContext(ctx)
		// 先校验jwt，再校验参数，jwt里有用户数据
		if err := ValidJwt(ctx); err != nil {
			resp = genHttpResp(beanerr.JwtTokenError.CloneWithError(err), nil)
			bizCtx.Log().Infof(GetRequestLog(ctx, nil, resp))
			ctx.JSON(0, resp)
			return
		}
		if len(req) != 0 {
			if bizErr = bizCtx.ParseRequest(req[0]); bizErr != nil {
				// 报错直接返回
				resp = genHttpResp(bizErr, nil)
				bizCtx.Log().Errorf(GetRequestLog(ctx, bizCtx.GetReqParam(), resp))
				ctx.JSON(0, genHttpResp(bizErr, nil))
				return
			}
		} else if len(req) > 1 {
			resp = genHttpResp(beanerr.ParamError.AppendMsg("too many param"), nil)
			bizCtx.Log().Errorf(GetRequestLog(ctx, bizCtx.GetReqParam(), resp))
			ctx.JSON(0, resp)
			return
		}
		data, bizErr := handler(bizCtx)
		resp = genHttpResp(bizErr, data)
		bizCtx.Log().Infof(GetRequestLog(ctx, bizCtx.GetReqParam(), resp))
		ctx.JSON(0, resp)
	}
}

func genHttpResp(err *beanerr.BizError, data interface{}) *HttpResponse {
	if err == nil {
		err = beanerr.NewBizError(0, "success")
	}
	r := new(HttpResponse)
	r.ErrCode = int32(err.ErrCode)
	r.Detail = err.Detail
	if data != nil {
		r.Data = data
	}
	r.Msg += err.Msg
	return r
}

func GetRequestLog(ctx *gin.Context, req, resp interface{}) string {
	return fmt.Sprintf("%s|%s|req: %v|resp: %v", ctx.Request.Method, ctx.Request.URL.Path, json_util.NoErrToJSON(req), json_util.NoErrToJSON(resp))
}
