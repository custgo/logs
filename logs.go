package logs

import (
	"io"
)

var originalDefaultLogger = NewLogger(&LogsConfig{
	Types: []string{"info", "warn", "error"},
	Files: map[string][]string{
		"STDOUT": []string{"info", "warn"},
		"STDERR": []string{"error"},
	},
})

var defaultLogger = originalDefaultLogger

func Restore() {
	defaultLogger = originalDefaultLogger
}

func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

func SetDefaultLoggerForConfig(conf *LogsConfig) {
	defaultLogger = NewLogger(conf)
}

func SetTimeFormat(fmt string) {
	defaultLogger.SetTimeFormat(fmt)
}

func SetPrefix(typeName string, prefix string) {
	defaultLogger.SetPrefix(typeName, prefix)
}

func SetTypes(types int) {
	defaultLogger.SetTypes(types)
}

func SetWriter(typeName string, writer io.Writer) {
	defaultLogger.SetWriter(typeName, writer)
}

func Debug(message string, args ...interface{}) {
	defaultLogger.Debug(message, args...)
}

func Info(message string, args ...interface{}) {
	defaultLogger.Info(message, args...)
}

func Warn(message string, args ...interface{}) {
	defaultLogger.Warn(message, args...)
}

func Error(message string, args ...interface{}) {
	defaultLogger.Error(message, args...)
}

func Fatal(message string, args ...interface{}) {
	defaultLogger.Fatal(message, args...)
}
