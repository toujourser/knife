package logger

import (
	"testing"
)

func init() {
	conf := &LogConf{
		Filename:   "./logs/t.log",
		MaxSize:    50,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
		LocalTime:  true,
		Level:      "debug",
		JsonFormat: false,
		StdOut:     true,
		NewLog:     true,
	}

	Init(conf)
}

func TestZaplog(t *testing.T) {

	Debug("test debug")
	Debugf("test %s", "debugf")
	Info("test info")
	Infof("test %+v", "infof")
	Warn("read failed")
	Warnf("%+v failed", "read")
	//Error("connection error")
	//Errorf("%s error", "connection")
	//Panic("process panic")
	//Panicf("%s panic", "process")
}
