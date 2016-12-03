package logs

import (
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	Debug("Debug:", "every thins goes right")
	Info("Info")
	Warn("Warning: ", "our base is under attack!")
	Error("Error: ", "div by zero")

	SetTypes(TYPE_ALL)
	SetPrefix("debug", "[TESTING]")
	SetTimeFormat("20060102150405")
	SetWriter("debug", os.Stdout)
	Debug("Debug:", "every thins goes right")
}

func TestLogfile(t *testing.T) {
	infofile := "/tmp/info.log"
	errorfile := "/tmp/error.log"
	conf := &LogsConfig{
		Types: []string{"debug", "warn", "error"},
		Files: map[string][]string{
			infofile:  []string{"debug", "info", "warn"},
			errorfile: []string{"error"},
		},
	}
	SetDefaultLoggerForConfig(conf)
	TestDefault(t)
	_, ierr := os.Stat(infofile)
	if os.IsNotExist(ierr) {
		t.Log("file info not exists")
		t.Fail()
	}
	_, eerr := os.Stat(errorfile)
	if os.IsNotExist(eerr) {
		t.Log("file error not exists")
		t.Fail()
	}
	if err := os.Remove(errorfile); err != nil {
		t.Log("remove ", errorfile, " faild: ", err)
	}
	if err := os.Remove(infofile); err != nil {
		t.Log("remove ", infofile, " faild: ", err)
	}
}

func TestGetExecPath(t *testing.T) {
	t.Log(GetExecPath())
}
