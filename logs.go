package logs

var originalDefaultLogger = NewLogger(&LogsConfig{
	Types: []string{"info", "warn", "error"},
	Files: map[string][]string{
		"STDOUT": []string{"info", "warn"},
		"STDERR": []string{"error"},
	},
})

var defaultLogger = originalDefaultLogger

func SetDefaultLogger(conf *LogsConfig) {
	defaultLogger = NewLogger(conf)
}

func SetTimeFormat(fmt string) {
	defaultLogger.SetTimeFormat(fmt)
}

func SetPrefix(typeName string, prefix string) {
	defaultLogger.SetPrefix(typeName, prefix)
}

func Debug(message string, args ...interface{}) {
	defaultLogger.Debug(message, args...)
}

func Info(message string, args ...interface{}) {
	defaultLogger.Info(message, args...)
}

func Warn(message string, args ...interface{}) {
	defaultLogger.Warn(message, args)
}

func Error(message string, args ...interface{}) {
	defaultLogger.Error(message, args...)
}

func Fatal(message string, args ...interface{}) {
	defaultLogger.Fatal(message, args...)
}
