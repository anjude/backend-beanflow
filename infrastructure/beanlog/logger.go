package beanlog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type BeanLogger struct {
	*logrus.Logger
}

var log = BeanLogger{
	Logger: logrus.New(),
}

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	// 设置自定义的日志格式化器
	log.SetFormatter(&BeanFormatter{})
}

// SetLevel 设定日志级别
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

func GetLogger(traceId string) *BeanLogger {
	newLog := log
	newLog.SetFormatter(&BeanFormatter{traceId: traceId})
	return &newLog
}

// Define logrus alias
var (
	Debugf = log.Debugf
	Infof  = log.Infof
	Warnf  = log.Warnf
	Errorf = log.Errorf
	Fatalf = log.Fatalf
	Panicf = log.Panicf
	Printf = log.Printf
	Info   = log.Info
	Debug  = log.Debug
	Error  = log.Error
)
