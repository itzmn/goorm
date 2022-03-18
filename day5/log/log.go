package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// 日志模块  自定义日志

// 定义两个日志输出实例
var (
	errorLogger = log.New(os.Stdout, "\033[31m[ERROR]\033[0m\t", log.LstdFlags|log.Lshortfile)
	infoLogger  = log.New(os.Stdout, "\033[34m[INFO]\033[0m\t", log.LstdFlags|log.Lshortfile)
	loggers     = []*log.Logger{errorLogger, infoLogger}
	mu          sync.Mutex

	// 定义日志打印的方法

	Error  = errorLogger.Println
	Errorf = errorLogger.Printf
	Info   = infoLogger.Println
	Infof  = infoLogger.Printf
)

// 日志级别
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLogger.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLogger.SetOutput(ioutil.Discard)
	}
}
