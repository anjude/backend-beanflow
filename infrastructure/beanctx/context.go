package beanctx

import (
	"net/http"
	"reflect"

	"github.com/anjude/backend-beanflow/infrastructure/beanerr"
	"github.com/anjude/backend-beanflow/infrastructure/beanlog"
	"github.com/anjude/backend-beanflow/infrastructure/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BizContext 扩展gin.Context
type BizContext struct {
	*gin.Context
	*metadata // 请求通用参数
	db        *gorm.DB
	log       *beanlog.BeanLogger
	request   interface{}
}

func NewContext(c *gin.Context) *BizContext {
	return &BizContext{
		Context:  c,
		request:  nil,
		log:      newLogger(c),
		metadata: newMetadata(c),
	}
}

func newLogger(c *gin.Context) *beanlog.BeanLogger {
	var traceId string
	if c.Request != nil {
		traceId = c.GetHeader("X-Trace-id")
	}
	if traceId == "" {
		traceId = generateRequestID()
	}
	return beanlog.GetLogger(traceId)
}

func generateRequestID() string {
	u := uuid.New()
	return u.String()
}

func (b *BizContext) ParseRequest(req interface{}) *beanerr.BizError {
	if req == nil {
		return nil
	}
	paramType := reflect.TypeOf(req)
	if paramType == nil {
		return nil
	}
	reqParam := reflect.New(paramType).Interface()

	var err error
	if b.Request.Method == http.MethodGet {
		err = b.ShouldBindQuery(reqParam)
	} else if b.Request.Method == http.MethodPost {
		err = b.ShouldBindJSON(reqParam)
	} else {
		b.log.Printf("uncompatible request method: %s", b.Request.Method)
		err = b.ShouldBind(reqParam)
	}
	if err != nil {
		return beanerr.ParamError.CloneWithError(err)
	}
	b.request = reflect.ValueOf(reqParam).Elem().Interface()
	return nil
}

func (b *BizContext) GetDb() *gorm.DB {
	// 开启事务时，使用传入的事务
	if b.db != nil {
		return b.db
	}
	return global.MysqlDB.GetDb()
}

func (b *BizContext) Log() *beanlog.BeanLogger {
	return b.log
}

// GetReqParam 获取解析&校验后的参数
func (b *BizContext) GetReqParam() interface{} {
	return b.request
}

// RunInTransaction 不支持嵌套事务
func (b *BizContext) RunInTransaction(fn func() error) error {
	if b.db != nil {
		return fn()
	}
	// 开始事务
	b.db = b.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			// 发生 panic，回滚事务
			b.db.Rollback()
			panic(r) // 继续抛出 panic
		}
	}()

	// 执行传入的函数，并将事务对象作为参数传递
	if err := fn(); err != nil {
		// 如果发生错误，回滚事务
		b.db.Rollback()
		return err
	}

	// 提交事务
	if err := b.db.Commit().Error; err != nil {
		b.db.Rollback()
		return err
	}

	return nil
}
