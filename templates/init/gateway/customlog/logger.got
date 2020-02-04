package customlog

import (
	"{{.ModuleName}}/appcontext"
	"{{.ModuleName}}/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger wrapper with two Logger methods
type Logger struct {
	//Logger only supports structured logging (less features more performance)
	Logger *zap.Logger
	//Sugar supports structured and printf-style APIs (less performance more features)
	Sugar *zap.SugaredLogger
}

func discoverLogLevel() zapcore.Level {
	switch config.Values.LogLevel {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "WARNING":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	case "PANIC":
		return zap.PanicLevel
	case "FATAL":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

//Get s the Current Logger component
func Get() Logger {
	return appcontext.Current.Get(appcontext.Logger).(Logger)
}

func init() {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(discoverLogLevel()),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	sugar := logger.Sugar()
	appcontext.Current.Add(appcontext.Logger, Logger{Logger: logger, Sugar: sugar})
}
