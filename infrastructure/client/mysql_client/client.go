package mysql_client

import (
	"fmt"
	"log"

	"github.com/anjude/backend-beanflow/infrastructure/config"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConn struct {
	db *gorm.DB
}

func NewMysqlConn(db *gorm.DB) *MySQLConn {
	return &MySQLConn{
		db: db,
	}
}

func (c *MySQLConn) GetDb() *gorm.DB {
	return c.db
}

func (c *MySQLConn) SetDb(db *gorm.DB) {
	c.db = db
}

func GetMySQLConn(conf config.DbConfig) (*gorm.DB, error) {
	db, err := initDB(conf)
	if err != nil {
		return nil, errors.WithMessage(err, "init mysql_client fail")
	}
	return db, nil
}

// initDB
func initDB(config config.DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Address, config.Database)))
	if err != nil {
		return nil, errors.WithMessage(err, "init db_client fail")
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, errors.WithMessage(err, "init db_client fail")
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Printf("mysql ping fail, err: %v", err)
		return nil, errors.WithMessage(err, "init db_client fail")
	}
	return db, err
}
