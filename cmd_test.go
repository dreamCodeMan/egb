package egb

import (
	"strings"
	"testing"
)

func TestExecCmd(t *testing.T) {
	stdout, stderr, err := ExecCmd("go", "help", "get")
	if err != nil {
		t.Errorf("ExecCmd:\n Expect => %v\n Got => %v\n", nil, err)
	} else if len(stderr) != 0 {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "", stderr)
	} else if !strings.HasPrefix(stdout, "usage: go get") {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "usage: go get", stdout)
	}
}

func TestExecCmdDir(t *testing.T) {
	out, errStr, err := ExecCmdDir("/Users/angelina-zf/Playground/golib/src/github.com/agelinazf/egb/concurrent-map", "du", "-sk")
	if err != nil {
		t.Error(err.Error())
		t.Error(errStr)
	}
	println(strings.Trim(strings.TrimSpace(out), "."))
}
