package boostrap

import (
	"fmt"
	"github.com/anjude/backend-beanflow/infrastructure/config"
	"github.com/anjude/backend-beanflow/infrastructure/constant"
	"github.com/anjude/backend-beanflow/infrastructure/enum"
	"github.com/anjude/backend-beanflow/infrastructure/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func InitConfig() error {
	v := viper.New()
	v.SetConfigName(getConfigName()) // name of config file (without extension)
	v.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	configPath := os.Getenv(enum.EkConfigPath.ToString())
	if configPath == "" {
		configPath = os.Getenv(strings.ToLower(enum.EkConfigPath.ToString()))
	}
	if configPath == "" {
		configPath = constant.DefaultConfigPath
	}
	v.AddConfigPath(configPath)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {
		return err
	}

	// 监听配置文件
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err = v.Unmarshal(&global.Conf); err != nil {
			// TODO-anjude 告警
			log.Fatalf("config file reload error: %v\n", err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.Conf); err != nil {
		return err
	}

	return nil
}

func getConfigName() string {
	env := config.GetEnv()
	if env == enum.LIVE {
		return "config"
	}
	return "config" + fmt.Sprintf(".%s", strings.ToLower(env.ToString()))
}
