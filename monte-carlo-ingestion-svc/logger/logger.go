package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger is the singleton instance of the Logger.
var logger *Logger

// Logger is a structure that holds the logger instance.
type Logger struct {
	Log *zap.Logger
}

// GetLogger returns the singleton instance of the logger.
func GetLogger() *Logger {
	return logger
}

// NewLogger creates a new logger instance.
func NewLogger() (*Logger, error) {
	return logger, nil
}

// init initializes the logger instance.
func init() {
	var zapLogger *zap.Logger
	var err error

	if os.Getenv("ENV") == "local" {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapLogger, err = config.Build()
	} else {
		zapLogger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	defer zapLogger.Sync()

	logger = &Logger{Log: zapLogger}
}

// Info logs the message with the fields at the info level.
func Info(msg string, fields ...zap.Field) {
	GetLogger().Log.Info(msg, fields...)
}

// Error logs the message with the fields at the error level.
func Error(msg string, fields ...zap.Field) {
	GetLogger().Log.Error(msg, fields...)
}

// Warn logs the message with the fields at the warn level.
func Warn(msg string, fields ...zap.Field) {
	GetLogger().Log.Warn(msg, fields...)
}

// Debug logs the message with the fields at the debug level.
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Log.Debug(msg, fields...)
}

// Fatal logs the message with the fields at the fatal level.
func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Log.Fatal(msg, fields...)
}
