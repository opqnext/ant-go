package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"	//使用Lumberjack进行日志切割归档
)

var log *zap.Logger

func InitLogger(logPath, logLevel string) error {

	lumberJackLogger := lumberjack.Logger{
		Filename: logPath,	//日志文件的位置
		MaxSize: 1024,		//在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 3,		//保留旧文件的最大个数
		MaxAge: 7,			//保留旧文件的最大天数
		Compress: true,		//是否压缩/归档旧文件
	}
	w := zapcore.AddSync(&lumberJackLogger)

	atom := zap.NewAtomicLevel()

	switch logLevel {
	case "debug":
		atom.SetLevel(zap.DebugLevel)
	case "info":
		atom.SetLevel(zap.InfoLevel)
	case "error":
		atom.SetLevel(zap.ErrorLevel)
	default:
		atom.SetLevel(zap.DebugLevel)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		atom)
	log = zap.New(core)

	return nil
}

func GetLogger() *zap.Logger  {
	return log
}
