package beanlog

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var _ logrus.Formatter = &BeanFormatter{}

type BeanFormatter struct {
	traceId string
}

// Format [time] [level] [file:line] [func] [msg] [data]
func (f BeanFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b strings.Builder

	// 添加时间戳
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05") + "|")

	// 添加日志级别
	b.WriteString(strings.ToUpper(entry.Level.String()) + "|")

	// 添加日志来源（文件名）
	// 提取最近一级的目录、文件名、函数名和行号
	if entry.Caller != nil {
		dirPath := extractDirectoryPath(entry.Caller.File)
		fileName := extractFileName(entry.Caller.File)
		funcName := extractFunctionName(entry.Caller.Function)
		line := entry.Caller.Line

		// 添加目录、文件名、函数名和行号
		b.WriteString(fmt.Sprintf("%s/%s:%d|%s|", dirPath, fileName, line, funcName))
	}

	// 添加请求 ID
	if f.traceId != "" {
		b.WriteString(fmt.Sprintf("%s|", f.traceId))
	}

	// 添加日志消息
	b.WriteString(entry.Message)

	// 添加换行符
	b.WriteString("\n\n")

	return []byte(b.String()), nil
}

// extractDirectoryPath 从完整文件路径中提取最近一级的目录路径
func extractDirectoryPath(path string) string {
	dir := filepath.Dir(path)
	return filepath.Base(dir)
}

// extractFileName 从完整文件路径中提取文件名
func extractFileName(path string) string {
	return filepath.Base(path)
}

// extractFunctionName 从完整函数名中提取函数名
func extractFunctionName(function string) string {
	// 函数名形式为包名.函数名，使用最后一个点进行分割
	parts := strings.Split(function, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return function
}
