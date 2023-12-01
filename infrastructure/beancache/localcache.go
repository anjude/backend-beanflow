package beancache

import (
	"time"

	"github.com/anjude/backend-beanflow/infrastructure/beanctx"

	"github.com/patrickmn/go-cache"
)

const (
	defaultExpiration = 5 * time.Minute
	cleanupInterval   = 1 * time.Minute
)

var localCache *cache.Cache

func InitLocalCache() {
	localCache = cache.New(defaultExpiration, cleanupInterval)
}

func GetInstance() *cache.Cache {
	if localCache == nil {
		InitLocalCache()
	}
	return localCache

}

func Set(ctx *beanctx.BizContext, key string, value interface{}, expiration time.Duration) {
	localCache.Set(key, value, expiration)
}

func Get(ctx *beanctx.BizContext, key string) (interface{}, bool) {
	return localCache.Get(key)
}

func GetString(key string) (string, bool) {
	if v, ok := localCache.Get(key); ok {
		return v.(string), ok
	}
	return "", false
}

func GetInt(key string) (int64, bool) {
	if v, ok := localCache.Get(key); ok {
		return v.(int64), ok
	}
	return 0, false
}

func GetBool(key string) (bool, bool) {
	if v, ok := localCache.Get(key); ok {
		return v.(bool), ok
	}
	return false, false
}
