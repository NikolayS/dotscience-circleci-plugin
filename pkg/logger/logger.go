package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLogLevel - default log level
var DefaultLogLevel = zap.InfoLevel

func getBaseLogConfig(level zapcore.Level) zapcore.EncoderConfig {

	switch level {
	case zap.DebugLevel:
		return zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Message",
			StacktraceKey:  "Stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	}

	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:  "Time",
		LevelKey: "Level",
		NameKey:  "Name",
		// CallerKey:      "Caller",
		MessageKey: "Message",
		// StacktraceKey:  "Stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// GetLoggerInstance - returns ready to use logger instance
func GetLoggerInstance(level zapcore.Level) *zap.Logger {

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(getBaseLogConfig(level)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		level,
	)
	return zap.New(core, zap.AddCaller())

}
