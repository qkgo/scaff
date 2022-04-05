package cfg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

func InitZapLog() *zap.SugaredLogger {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.Config{
		DisableStacktrace: false,
		DisableCaller:     false,
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		//Level:             zap.NewAtomicLevelAt(zap.WarnLevel),
		//Level:             zap.NewAtomicLevelAt(zap.ErrorLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "lv",
			NameKey:        "logger",
			CallerKey:      "line",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}

func ZError(log zap.SugaredLogger, logInfo ...string) {
	funcName, file, line, _ := runtime.Caller(0)
	log.Errorf(file, funcName, line, logInfo)
}
