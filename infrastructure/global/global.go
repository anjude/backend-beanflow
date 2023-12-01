package global

import (
	"github.com/anjude/backend-beanflow/infrastructure/client/mysql_client"
	"github.com/anjude/backend-beanflow/infrastructure/config"
)

var (
	MysqlDB *mysql_client.MySQLConn
	Conf    *config.Config
)
