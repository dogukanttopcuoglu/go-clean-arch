package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(message string)
	Warn(err error, message string)
	Error(err error, message string)
	Sync() error
}

type zapLogger struct {
	logger *zap.Logger
}

type Options struct {
	Environment string
	Level       string
	ServiceName string
}

func New(opts Options) (Logger, error) {
	level := parseLevel(opts.Level)

	environment := strings.ToLower(strings.TrimSpace(opts.Environment))

	isDevelopment := environment == "development" ||
		environment == "dev" ||
		environment == "local"

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	encoding := "json"

	if isDevelopment {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.MessageKey = "message"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		encoding = "console"
	}

	initialFields := map[string]interface{}{
		"pid": os.Getpid(),
	}

	if opts.ServiceName != "" {
		initialFields["service"] = opts.ServiceName
	}

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       isDevelopment,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: initialFields,
	}

	zapLog, err := cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		logger: zapLog,
	}, nil
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return zapcore.DebugLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *zapLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *zapLogger) Warn(err error, message string) {
	l.logger.Warn(message, zap.Error(err))
}

func (l *zapLogger) Error(err error, message string) {
	l.logger.Error(message, zap.Error(err))
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}
