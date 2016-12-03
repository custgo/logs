package logs

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	NOLOG = 0
	DEBUG = 1
	INFO  = 2
	WARN  = 4
	ERROR = 8
	ALL   = 15
)

type LogsConfig struct {
	Types []string            `json:types`
	Files map[string][]string `json:files`
}

// 返回 LogsConfig 中配置的日志类型
// 例如设置了 ERROR + INFO，则返回 10 （ERROR | INFO）
func (conf *LogsConfig) getTypes() int {
	lg := 0
	if 0 == len(conf.Types) {
		return lg
	}
	for _, ln := range conf.Types {
		lg |= getTypeByName(ln)
	}
	return lg
}

// 解析 LogsConfig 的 files，生成与日志类型关联的 io.Writer
// 例如 {ERROR: [os.Stderr, os.Stdout]}
func (conf *LogsConfig) getWriters() map[int][]io.Writer {
	ret := make(map[int][]io.Writer)
	for fileName, types := range conf.Files {
		writer := getWriterByName(fileName)
		for _, logType := range types {
			itype := getTypeByName(logType)
			if _, exists := ret[itype]; !exists {
				ret[itype] = make([]io.Writer, 0)
			}
			ret[itype] = append(ret[itype], writer)
		}
	}
	return ret
}

var writers_map = make(map[string]io.Writer)

func getWriterByName(name string) io.Writer {
	if strings.HasPrefix(name, "{AppPath}") {
		name = getExecPath() + name[9:]
	}
	writer, exists := writers_map[name]
	if exists {
		return writer
	}
	if "STDOUT" == strings.ToUpper(name) {
		return os.Stdout
	}
	if "STDERR" == strings.ToUpper(name) {
		return os.Stderr
	}
	logf, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		log.Fatal("Error! Can not Open Log File:", err)
	}
	writers_map[name] = logf
	return logf
}

func getTypeByName(name string) int {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	}
	return NOLOG
}

var execPath string

func getExecPath() string {
	if "" == execPath {
		execFile, _ := exec.LookPath(os.Args[0])
		execPath = filepath.Dir(execFile)
	}
	return execPath
}
