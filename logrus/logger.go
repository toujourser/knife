package logrus

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	level := strings.ToUpper(entry.Level.String())
	file := entry.Data["file"]
	line := entry.Data["line"]
	message := entry.Message
	var fields string
	for key, val := range entry.Data {
		if key == "file" || key == "line" || key == "func" {
			continue
		}
		fields += fmt.Sprintf("%s=%v ", key, val)
	}

	logLine := fmt.Sprintf("%s [%s] %v:%d %s %s\n", timestamp, level, file, line, fields, message)
	return []byte(logLine), nil
}

func Init(logPath, logName string, maxAge, rotationTime int) {
	logger = logrus.New()
	logger.SetFormatter(&CustomFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(newRotateHook(path.Join(logPath, logName), time.Duration(maxAge)*24*time.Hour, time.Duration(rotationTime)))

}

func newRotateHook(baseLogPath string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logger.Errorf("config local file system Logger error. %+v", errors.WithStack(err))
	}
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &CustomFormatter{})
}

func Debug(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	logger.WithFields(getCallerInfo()).Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.WithFields(getCallerInfo()).Panicf(format, args...)
}

func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}

func WithField(key string, val interface{}) *logrus.Entry {
	return logger.WithFields(getCallerInfo()).WithField(key, val)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(getCallerInfo()).WithFields(fields)
}

func getCallerInfo() logrus.Fields {
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	_, fileName := filepath.Split(file)
	_, funcName = filepath.Split(funcName)

	return logrus.Fields{
		"file": fileName,
		"func": funcName,
		"line": line,
	}
}
