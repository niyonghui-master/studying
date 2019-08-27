package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zbindenren/logrus_mail"
	"time"
)

func main() {
	logger := logrus.New()
	hook, err := logrus_mail.NewMailAuthHook("logrus_mail", "smtp.gmail.com", 587, "18516383338@163.com", "1243600515@qq.com", "18516383338@163.com", "123456")

	if err == nil {
		logger.Hooks.Add(hook)
	}

	// 生成*Entry
	var filename = "123.txt"
	contextLogger := logger.WithFields(logrus.Fields{
		"file":    filename,
		"context": "GG",
	})
	// 设置时间戳和message
	contextLogger.Time = time.Now()
	contextLogger.Message = "这是一个hook发来的邮件"
	// 只能发送Error, Fatal, panic级别的log
	contextLogger.Level = logrus.ErrorLevel

	// 使用Fire发送，包含时间戳，message
	hook.Fire(contextLogger)
}

// 默认的Logger在并发写入的时候有mutex保护，其在调用hooks和写入logs时被启用。如果你认为此锁是没有必要的，可以添加logger.SetNoLock()来让锁失效。