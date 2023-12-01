package boostrap

import "github.com/anjude/backend-beanflow/infrastructure/beanlog"

func InitLogger() error {
	beanlog.InitLogger()
	return nil
}
