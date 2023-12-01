package main

import (
	"fmt"
	"log"

	"github.com/anjude/backend-beanflow/infrastructure/boostrap"
	"github.com/anjude/backend-beanflow/infrastructure/global"
	"github.com/anjude/backend-beanflow/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	// 新建一个空的gin实例
	engine := gin.New()

	// 初始化配置
	err := boostrap.InitConfig()
	if err != nil {
		log.Fatalf("init config fail, err: %v", err)
	}

	// 初始化各个模块
	initModule := []func() error{
		boostrap.InitLogger,
		boostrap.InitMysql,
		boostrap.InitRedis,
		boostrap.InitLocalCache,
	}
	for _, f := range initModule {
		if err = f(); err != nil {
			log.Fatalf("init module fail, err: %v", err)
		}
	}

	// 注册中间件
	engine.Use(gin.Recovery())

	// 初始化api服务的路由
	interfaces.NewApiService().RegisterRouter(engine)

	// 启动服务
	err = engine.Run(fmt.Sprintf(":%v", global.Conf.App.Port))
	if err != nil {
		log.Fatalf("init server fail, err: %v", err)
	}
}
