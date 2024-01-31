package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log *zap.SugaredLogger
)

type LogConf struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	LocalTime  bool
	Level      string
	JsonFormat bool
	StdOut     bool
	NewLog     bool
}

func Init(config *LogConf) {
	var syncWriters []zapcore.WriteSyncer
	lumberjackCfg := &lumberjack.Logger{
		Filename:   config.Filename,   // ⽇志⽂件路径
		MaxSize:    config.MaxSize,    // 单位为MB,默认为512MB
		MaxAge:     config.MaxAge,     // 文件最多保存多少天
		MaxBackups: config.MaxBackups, // 最多备份几个
		LocalTime:  config.LocalTime,  // 采用本地时间
		Compress:   config.Compress,   // 是否压缩日志
	}

	if config.NewLog {
		_ = lumberjackCfg.Rotate()
	}
	syncWriters = append(syncWriters, zapcore.AddSync(lumberjackCfg))

	if config.StdOut {
		syncWriters = append(syncWriters, zapcore.AddSync(os.Stdout))
	}

	syncWriter := zapcore.NewMultiWriteSyncer(syncWriters...)

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.0000"))
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		s := level.CapitalString()
		enc.AppendString(s)
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		s := caller.TrimmedPath()
		enc.AppendString("(" + caller.Function + ") " + s)
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller",
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeLevel:    customLevelEncoder,  // 小写编码器
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		//EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色打印
	}

	// 设置日志级别
	atomicLevel := getLogLevel(config.Level)

	// json还是字符
	var encode zapcore.Encoder
	if config.JsonFormat {
		encode = zapcore.NewJSONEncoder(encoderConf)
	} else {
		encode = zapcore.NewConsoleEncoder(encoderConf)
	}

	core := zapcore.NewCore(
		encode,      // 编码器配置
		syncWriter,  // 打印到文件,打印到控制台需要添加 zapcore.AddSync(os.Stdout)
		atomicLevel, // 日志级别
	)

	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func getLogLevel(lvl string) (atomicLevel zapcore.Level) {
	switch lvl {
	case "debug":
		atomicLevel = zap.DebugLevel
	case "info":
		atomicLevel = zap.InfoLevel
	case "warn":
		atomicLevel = zap.WarnLevel
	case "error":
		atomicLevel = zap.ErrorLevel
	case "fatal":
		atomicLevel = zap.FatalLevel
	case "panic":
		atomicLevel = zap.PanicLevel
	default:
		atomicLevel = zap.InfoLevel
	}
	return
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Sugar() {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	consoleWriter := zapcore.Lock(os.Stdout)
	fileWriter, _ := os.OpenFile("/tmp/hardware_info/result_analysis.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(consoleWriter), zapcore.DebugLevel),
	)
	_ = zap.New(core).WithOptions(zap.AddCaller()).Sugar()
}
