package logger

import (
	"context"

	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
func init() {
	fmt.Println("init to logger---------------------")
	LoggerInit("log.log", zapcore.InfoLevel)
	fmt.Println("init to logger---------------------")
}

var logObject *zap.Logger

func Log(data ...context.Context) *zap.Logger {
	if data != nil {
		//ctx := data[0]
		return logObject.With(zap.Any("Token", "ctx.Value(config.TOKEN)"))
	} else {
		return logObject
	}
}

func LoggerInit(logFilePath string, level zapcore.Level) {
	var (
		err error
	)
	//fmt.Println("LOGGER INIT started")
	//logFile, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.FunctionKey = "f"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	cfg.EncoderConfig.ConsoleSeparator = " "
	cfg.EncoderConfig.EncodeCaller = MyCaller
	cfg.Encoding = "console"
	cfg.DisableStacktrace = true // no need of line no 62 with options
	// if logFilePath != "" {
	// 	var paths []string
	// 	paths = append(paths, logFilePath)
	// 	cfg.OutputPaths = paths
	// }
	//TODO: Error handling
	logObject, err = cfg.Build()
	if err != nil {
		fmt.Println("failed to create custom production logger , Exiting system", err)
		os.Exit(0)
	} else if logObject == nil {
		logObject, err = zap.NewProduction()
		logObject.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.InfoLevel))
		if err != nil {
			fmt.Println("failed to create production logger , Exiting system", err)
			os.Exit(0)
		}
		fmt.Println("Failed to create custom logger, creating production logger")
	} else {
		//logObject.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
		//fmt.Println("custom production logger created")
	}

	Log().Info("Logger init successfully")
}

// time logger
func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}
