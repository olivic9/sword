package exceptions

import (
	"fmt"
	"runtime"
)

type Error interface {
	Error() string
	Code() int
}

type HandledError struct {
	s string
	c int
}

func (e *HandledError) Error() string {
	return e.s
}

func (e *HandledError) Code() int {
	return e.c
}

func New(s string, c int) Error {
	return &HandledError{s, c}
}

func NewInternalHandledError(error string) Error {
	pc, filename, line, _ := runtime.Caller(1)
	return New(fmt.Sprintf("[error] function: %s  file: %s:%d error: %v", runtime.FuncForPC(pc).Name(), filename, line, error), 500)
}
