package logging

import (
	"io"
	"log"
)

// 日志等级
const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

type NLog struct {
	Level  int // 日志输出的等级
	L      *log.Logger
	Module string
}

func NewNLog(out io.Writer, level int, module string, flag int) *NLog {
	logger := log.New(out, "", flag)
	return &NLog{
		Level:  level,
		L:      logger,
		Module: module,
	}
}

// 获取日志对应的前缀
func getLevelPrefix(level int) string {
	switch level {
	case LevelDebug:
		return "[D]"
	case LevelError:
		return "[E]"
	case LevelInfo:
		return "[I]"
	case LevelWarning:
		return "[W]"
	case LevelFatal:
		return "[F]"
	}
	return ""
}

// 格式化Module输出
func (N *NLog) formatModule() []byte {
	return []byte(" [" + N.Module + "] ")
}

func (N *NLog) Debug(v ...interface{}) {
	if N.Level < LevelDebug {
		return
	}

	EPrefix := getLevelPrefix(LevelDebug)
	LogPrefix := append([]byte(EPrefix), N.formatModule()...)
	N.L.SetPrefix(string(LogPrefix))
	N.L.Println(v...)
}

func (N *NLog) Info(v ...interface{}) {
	if N.Level < LevelInfo {
		return
	}

	EPrefix := getLevelPrefix(LevelInfo)
	LogPrefix := append([]byte(EPrefix), N.formatModule()...)
	N.L.SetPrefix(string(LogPrefix))
	N.L.Println(v...)
}

func (N *NLog) Warning(v ...interface{}) {
	if N.Level < LevelWarning {
		return
	}

	EPrefix := getLevelPrefix(LevelWarning)
	LogPrefix := append([]byte(EPrefix), N.formatModule()...)
	N.L.SetPrefix(string(LogPrefix))
	N.L.Println(v...)
}

func (N *NLog) Error(v ...interface{}) {
	if N.Level < LevelError {
		return
	}

	EPrefix := getLevelPrefix(LevelError)
	LogPrefix := append([]byte(EPrefix), N.formatModule()...)
	N.L.SetPrefix(string(LogPrefix))
	N.L.Println(v...)
}

func (N *NLog) Fatal(v ...interface{}) {
	if N.Level < LevelFatal {
		return
	}

	EPrefix := getLevelPrefix(LevelFatal)
	LogPrefix := append([]byte(EPrefix), N.formatModule()...)
	N.L.SetPrefix(string(LogPrefix))
	N.L.Fatalln(v...)
}
