package initial

import (
	"gateway_server/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

// optionFunc wraps a func so it satisfies the Option interface.
// An Option configures a Logger.
type Config interface {
	apply(*loggerOptions)
}

type loggerOptions struct {
	zapOptions []zap.Option
}

// optionFunc wraps a func so it satisfies the Option interface.
type configFunc func(logger *loggerOptions)

func (f configFunc) apply(log *loggerOptions) {
	f(log)
}

func WrapOptions(opts []zap.Option) Config {
	return configFunc(func(l *loggerOptions) {
		l.zapOptions = opts
	})
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func InitAllZap(opts ...configFunc) {
	config := &loggerOptions{}
	for _, o := range opts {
		o.apply(config)
	}
	options := []zap.Option{}
	options = append(options, config.zapOptions...)
	tee := zapcore.NewTee(NewErrorCore(), NewWarnCore())
	logger := zap.New(tee, options...)
	zap.ReplaceGlobals(logger)
	zap.S().Infof("[Navi Gateway]	Init Zap Log success")
}

func NewErrorCore() zapcore.Core {
	return zapcore.NewCore(
		GetEncoder(),
		NewErrorWrite(),
		zap.ErrorLevel,
	)
}

func NewWarnCore() zapcore.Core {
	return zapcore.NewCore(
		GetEncoder(),
		NewWarnWrite(),
		zap.WarnLevel,
	)
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func NewErrorWrite() zapcore.WriteSyncer {
	path := GetCurrentPath()
	cfg := global.DebugFullConfig.ZapConfig
	l := &lumberjack.Logger{
		Filename:   path + cfg.ErrorPath + "/error.log",
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackup,
		Compress:   false,
	}
	return zapcore.AddSync(l)
}

func NewWarnWrite() zapcore.WriteSyncer {
	path := GetCurrentPath()
	cfg := global.DebugFullConfig.ZapConfig
	l := &lumberjack.Logger{
		Filename:   path + cfg.OtherPath + "/info.log",
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackup,
		Compress:   false,
	}
	return zapcore.AddSync(l)
}
