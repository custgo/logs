package logs

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu   sync.Mutex
	conf *LogsConfig

	types   int
	writers map[int][]io.Writer

	timeFmt  string
	prefixes map[int]string

	buf []byte
}

func NewLogger(conf *LogsConfig) *Logger {
	return &Logger{
		conf:    conf,
		types:   conf.getTypes(),
		writers: conf.getWriters(),
		timeFmt: "2006-01-02 15:04:05",
		prefixes: map[int]string{
			DEBUG: "[Debug]",
			INFO:  "[Info]",
			WARN:  "[Warn]",
			ERROR: "[Error]",
		},
	}
}

func (l *Logger) SetTimeFormat(fmt string) {
	l.mu.Lock()
	l.timeFmt = fmt
	l.mu.Unlock()
}

func (l *Logger) SetPrefix(typeName string, prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	itype := getTypeByName(typeName)
	if itype > NOLOG {
		l.prefixes[itype] = prefix
	}
}

func (l *Logger) SetTypes(types int) {
	l.mu.Lock()
	l.types = types
	l.mu.Unlock()
}

func (l *Logger) SetWriter(typeName string, writer io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	itype := getTypeByName(typeName)
	if itype <= NOLOG {
		return
	}
	writers, exists := l.writers[itype]
	if !exists {
		writers = make([]io.Writer, 0)
		l.writers[itype] = append(writers, writer)
		return
	}
	for _, w := range writers {
		if w == writer {
			exists = true
			break
		}
	}
	if !exists {
		l.writers[itype] = append(writers, writer)
	}
}

func (l *Logger) Write(itype int, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.types&itype == 0 {
		return
	}

	writers, exists := l.writers[itype]
	if !exists || 0 == len(writers) {
		return
	}

	l.buf = l.buf[:0]
	now := time.Now().Format(l.timeFmt)
	prefix, ok := l.prefixes[itype]
	if !ok {
		prefix = ""
	}
	l.buf = append(l.buf, now...)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, prefix...)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, message...)
	if len(message) == 0 || message[len(message)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	for _, w := range writers {
		w.Write(l.buf)
	}
}

func (l *Logger) WriteTypes(types int, message string) {
	for i := 1; i <= ALL; i <<= 1 {
		if i&types == i {
			l.Write(i, message)
		}
	}
}

func (l *Logger) WriteArgs(itype int, message string, args ...interface{}) {
	var msg string
	if len(args) > 0 {
		msg = message + fmt.Sprint(args...)
	} else {
		msg = message
	}
	l.Write(itype, msg)
}

func (l *Logger) WriteTypesArgs(types int, message string, args ...interface{}) {
	var msg string
	if len(args) > 0 {
		msg = message + fmt.Sprint(args...)
	} else {
		msg = message
	}
	l.WriteTypes(types, msg)
}

func (l *Logger) Debug(message string, args ...interface{}) {
	l.WriteArgs(DEBUG, message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.WriteArgs(INFO, message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.WriteArgs(WARN, message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.WriteArgs(ERROR, message, args...)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	l.Error(message, args...)
	os.Exit(1)
}
