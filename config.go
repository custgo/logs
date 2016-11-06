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

func getWriterByName(name string) io.Writer {
	if "STDOUT" == strings.ToUpper(name) {
		return os.Stdout
	}
	if "STDERR" == strings.ToUpper(name) {
		return os.Stderr
	}
	if strings.HasPrefix(name, "{AppPath}") {
		name = getExecPath() + name[9:]
	}
	logf, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		log.Fatal("Error! Can not Open Log File:", err)
	}
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
