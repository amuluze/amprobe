// Package logger
// Date: 2024/3/6 10:55
// Author: Amu
// Description:
package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

var defaultLevel slog.LevelVar
var defaultLogger *slog.Logger

type Logger struct {
	logger *slog.Logger
}

func init() {
	defaultLevel.Set(slog.LevelError)
	options := slog.HandlerOptions{
		AddSource: true,
		Level:     &defaultLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}
	defaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, &options))
}

func Debug(debugMsg string) {
	defaultLogger.Debug(debugMsg)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debug(fmt.Sprintf(format, args...))
}

func Info(infoMsg string) {
	defaultLogger.Info(infoMsg)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Info(fmt.Sprintf(format, args...))
}

func Warn(warnMsg string) {
	defaultLogger.Warn(warnMsg)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warn(fmt.Sprintf(format, args...))
}

func Error(errMsg string) {
	defaultLogger.Error(errMsg)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Error(fmt.Sprintf(format, args...))
}

func SetDebugLevel() {
	defaultLevel.Set(slog.LevelDebug)
}

func SetInfoLevel() {
	defaultLevel.Set(slog.LevelInfo)
}

func SetWarnLevel() {
	defaultLevel.Set(slog.LevelWarn)
}

func SetErrorLevel() {
	defaultLevel.Set(slog.LevelError)
}
