package beego_log

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

type MyLogger struct {
	log    *logs.BeeLogger
	fields map[string]interface{} // 扩展的 fields
}

func NewLogger(level int, logFile string) (l *MyLogger, err error) {
	configs := fmt.Sprintf(`{"filename":"%s"}`, logFile)
	logger := logs.NewLogger()
	_ = logger.SetLogger(logs.AdapterConsole)
	err = logger.SetLogger(logs.AdapterMultiFile, configs)
	if err != nil {
		return
	}
	logger.SetLevel(level)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(4)

	return &MyLogger{
		log:    logger,
		fields: make(map[string]interface{}),
	}, nil
}

// WithField 添加扩展字段
func (l *MyLogger) WithField(key string, value interface{}) *MyLogger {
	l.fields[key] = value
	return l
}

// WithFields 添加扩展字段（多个）
func (l *MyLogger) WithFields(fields map[string]interface{}) *MyLogger {
	for key, value := range fields {
		l.fields[key] = value
	}
	return l
}

func (l *MyLogger) print(level int, v ...interface{}) {
	message := fmt.Sprint(v...)

	// 拼接扩展字段
	for key, value := range l.fields {
		message += fmt.Sprintf(" %s=%v", key, value)
	}

	switch level {
	case logs.LevelEmergency:
		l.log.Emergency(message)
	case logs.LevelAlert:
		l.log.Alert(message)
	case logs.LevelCritical:
		l.log.Critical(message)
	case logs.LevelError:
		l.log.Error(message)
	case logs.LevelWarning:
		l.log.Warning(message)
	case logs.LevelNotice:
		l.log.Notice(message)
	case logs.LevelInformational:
		l.log.Informational(message)
	case logs.LevelDebug:
		l.log.Debug(message)
	}
}

// 封装所有 Level 方法
func (l *MyLogger) Emergency(v ...interface{}) {
	l.print(logs.LevelEmergency, v...)
}

func (l *MyLogger) Alert(v ...interface{}) {
	l.print(logs.LevelAlert, v...)
}

func (l *MyLogger) Critical(v ...interface{}) {
	l.print(logs.LevelCritical, v...)
}

func (l *MyLogger) Error(v ...interface{}) {
	l.print(logs.LevelError, v...)
}

func (l *MyLogger) Warn(v ...interface{}) {
	l.print(logs.LevelWarning, v...)
}

func (l *MyLogger) Notice(v ...interface{}) {
	l.print(logs.LevelNotice, v...)
}

func (l *MyLogger) Info(v ...interface{}) {
	l.print(logs.LevelInformational, v...)
}

func (l *MyLogger) Debug(v ...interface{}) {
	l.print(logs.LevelDebug, v...)
}
