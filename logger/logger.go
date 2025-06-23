package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var (
	mylog = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func InitLogger(format, level string) {
	var l slog.Level
	switch strings.ToUpper(level) {
	case "DEBUG":
		l = slog.LevelDebug
	case "INFO":
		l = slog.LevelInfo
	case "WARN":
		l = slog.LevelWarn
	case "ERROR", "FATAL":
		l = slog.LevelError
	default:
		l = slog.LevelDebug
	}
	slog.SetLogLoggerLevel(l)

	var handler slog.Handler
	switch strings.ToUpper(format) {
	case "JSON":
		handler = slog.NewJSONHandler(os.Stdout, nil)
	case "TEXT":
		handler = slog.NewTextHandler(os.Stdout, nil)
	default:
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}
	mylog = slog.New(handler)
}

func Debug(msg string, args ...any) {
	mylog.Debug(msg, args...)
}

func Debugf(format string, args ...any) {
	mylog.Debug(fmt.Sprintf(format, args...))
}

func Info(msg string, args ...any) {
	mylog.Info(msg, args...)
}

func Infof(format string, args ...any) {
	mylog.Info(fmt.Sprintf(format, args...))
}

func Warn(msg string, args ...any) {
	mylog.Warn(msg, args...)
}

func Warnf(format string, args ...any) {
	mylog.Warn(fmt.Sprintf(format, args...))
}

func Error(msg string, args ...any) {
	mylog.Error(msg, args...)
}

func Errorf(format string, args ...any) {
	mylog.Error(fmt.Sprintf(format, args...))
}

func Fatal(msg string, args ...any) {
	mylog.Error(msg, args...)
	panic(msg)
}

func Fatalf(format string, args ...any) {
	mylog.Error(fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}
