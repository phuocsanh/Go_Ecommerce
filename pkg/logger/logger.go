package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
type LoggerZap struct {
	*zap.Logger
}
func NewLogger() *LoggerZap{
	logLevel := "debug"
	// debug->info->warning->error->fatal->panic
	var level zapcore.Level
	switch logLevel{
		case "debug":
            level = zapcore.DebugLevel
        case "info":
            level = zapcore.InfoLevel
        case "warning":
            level = zapcore.WarnLevel
        case "error":
            level = zapcore.ErrorLevel
        case "fatal":
            level = zapcore.FatalLevel
        case "panic":
            level = zapcore.PanicLevel
        default:
            level = zapcore.InfoLevel
	}
	endCoder:= getEndCoderLog()
	hook := lumberjack.Logger{
		Filename:   "./storages/logs/dev.xxx.log",
    MaxSize:    500, // megabytes
    MaxBackups: 3,
    MaxAge:     28, //days
    Compress:   true, // disabled by default
	}
	core :=  zapcore.NewCore(endCoder,zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(&hook)), level )
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core,zap.AddCaller(),zap.AddStacktrace(zap.ErrorLevel))}

}

func getEndCoderLog() zapcore.Encoder {

	endCodeConfig := zap.NewProductionEncoderConfig()
	endCodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder // timetamps 1721898691.973886 ->  2024-07-25T16:11:31.956+0700  
	endCodeConfig.TimeKey = "time" // ts -> time
	endCodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // info -> INFO
	endCodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // cli/main.log.go:19
	return zapcore.NewJSONEncoder(endCodeConfig)

}
