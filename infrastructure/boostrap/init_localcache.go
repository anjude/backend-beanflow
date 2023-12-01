package boostrap

import "github.com/anjude/backend-beanflow/infrastructure/beancache"

func InitLocalCache() error {
	beancache.InitLocalCache()
	return nil
}
