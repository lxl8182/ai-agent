package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func InitLogger(level string, logFile string) error {
	var config zap.Config
	
	if level == "debug" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	// 设置日志级别
	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// 配置输出
	config.OutputPaths = []string{"stdout"}
	if logFile != "" {
		config.OutputPaths = append(config.OutputPaths, logFile)
	}

	var err error
	log, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

func Info(msg string, fields ...zap.Field) {
	if log != nil {
		log.Info(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if log != nil {
		log.Error(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if log != nil {
		log.Debug(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if log != nil {
		log.Warn(msg, fields...)
	}
}
