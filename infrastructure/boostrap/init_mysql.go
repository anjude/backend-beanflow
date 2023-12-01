package boostrap

import (
	"github.com/anjude/backend-beanflow/infrastructure/client/mysql_client"
	"github.com/anjude/backend-beanflow/infrastructure/global"
)

func InitMysql() error {
	db, err := mysql_client.GetMySQLConn(global.Conf.DbConfig)
	if err != nil {
		return err
	}
	global.MysqlDB = mysql_client.NewMysqlConn(db)
	return nil
}
