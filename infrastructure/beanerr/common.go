package beanerr

var (
	ParamError    = NewBizError(10001, "param error")
	InternalError = NewBizError(10002, "internal error")
	NoPermission  = NewBizError(10003, "no permission")
	JwtTokenError = NewBizError(10004, "jwt token err")
	DBError       = NewBizError(10005, "db error")
	ExternalError = NewBizError(10006, "external error")
)
