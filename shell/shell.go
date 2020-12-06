package shell

import (
	"os/exec"
	"strconv"
)

const (
	exit = ";echo $?"
	sync = "sync;sync"
)

// Run execute shell command and return "echo $?" int value.
func Run(command string) int {
	var ret int
	if len(command) > 0 {
		cmd := exec.Command("sh", "-c", command + exit)
		stdout, _ := cmd.CombinedOutput()
		ret, _ = strconv.Atoi(string(stdout))
	} else {
		ret = -1
	}
	return ret
}

// RunRead execute shell command and return output string.
func RunRead(command string) string {
	cmd := exec.Command("sh", "-c", command)
	stdout, _ := cmd.CombinedOutput()
	ret := string(stdout)
	return ret
}

// Sync execute sync command in shell.
func Sync() {
	cmd := exec.Command("sh", "-c", sync)
	cmd.Run()
}