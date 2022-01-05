package lumerinlib

import (
	"runtime"
	"strconv"
	"strings"
)

func BoilerPlateLibFunc(msg string) string {
	return msg
}

func FileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "FileLine() failed"
	}

	f := strings.Split(file, "/")

	lineno := strconv.Itoa(line)

	return "[" + f[len(f)-1] + ":" + lineno + "]:"
}

func Funcname() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "TheUnknownFunction()"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "TheUnknownFunction()"
	}

	f := strings.Split(fn.Name(), "/")

	return f[len(f)-1]
}

func Errtrace() string {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "file?[0]:func?"
	}

	lineno := strconv.Itoa(line)

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file + "[" + lineno + "]:func?"
	}

	return file + "[" + lineno + "]:" + fn.Name()
}