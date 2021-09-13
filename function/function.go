package function

import (
	"runtime"
)

func Line() int {
	_, _, line, _ := runtime.Caller(0)
	return line
}

func Name() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func Path() string {
	_, path, _, _ := runtime.Caller(0)
	return path
}