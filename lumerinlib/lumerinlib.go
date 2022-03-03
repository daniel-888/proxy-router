package lumerinlib

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type ConcurrentMap struct {
	sync.RWMutex
	M map[string]interface{}
}

func (r *ConcurrentMap) Get(key string) interface{} {
	r.RLock()
	defer r.RUnlock()
	return r.M[key]
}

func (r *ConcurrentMap) GetAll() (vals []interface{}) {
	r.RLock()
	defer r.RUnlock()
	for _, v := range r.M {
		vals = append(vals, v)
	}
	return vals
}

func (r *ConcurrentMap) Set(key string, val interface{}) {
	r.Lock()
	defer r.Unlock()
	r.M[key] = val
}

func (r *ConcurrentMap) Exists(key string) bool {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.M[key]
	return ok
}

func (r *ConcurrentMap) Delete(key string) {
	r.Lock()
	defer r.Unlock()
	delete(r.M, key)
}

//
//
//
func BoilerPlateLibFunc(msg string) string {
	return msg
}

//
//
//
func FileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "FileLine() failed"
	}

	f := strings.Split(file, "/")

	lineno := strconv.Itoa(line)

	return "[" + f[len(f)-1] + ":" + lineno + "]:"
}

//
//
//
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

//
//
//
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

//
//
//
func PanicHere(text ...string) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Well this is unexpected...")
	}

	f := strings.Split(file, "/")

	lineno := strconv.Itoa(line)

	panic(fmt.Sprintf("[%s:%s]:%s", f[len(f)-1], lineno, text[0]))
}
