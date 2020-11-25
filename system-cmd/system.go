package system

import (
	"os/exec"
	"strconv"
)

const (
	EXITSTR = ";echo $?"
	SYNCSTR = "sync;sync"
)

func Run(command string) int {
	var ret int
	if len(command) > 0 {
		cmd := exec.Command("sh", "-c", command + EXITSTR)
		stdout, _ := cmd.CombinedOutput()
		ret, _ = strconv.Atoi(string(stdout))
	} else {
		ret = -1
	}
	return ret
}

func RunRead(command string) string {
	cmd := exec.Command("sh", "-c", command)
	stdout, _ := cmd.CombinedOutput()
	ret := string(stdout)
	return ret
}

func Sync() {
	cmd := exec.Command("sh", "-c", SYNCSTR)
	cmd.Run()
}