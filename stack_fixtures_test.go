package wrapperr_test

import (
	"github.com/dc0d/wrapperr"
)

func fn1() error {
	return wrapperr.WithStack(errRootCause)
}

func fn2() error {
	return fn1()
}

func fn3() error {
	return wrapperr.WithStack(errRootCause, "message 3")
}

func fn4() error {
	return fn3()
}

func fn5() error {
	return wrapperr.WithStack(fn4(), "message 5")
}

func fn6() error {
	return fn5()
}

func fn7() error {
	return fn3()
}

func fn8() error {
	return wrapperr.WithStack(fn7(), emptyAnnotation)
}

func fn9() error {
	return fn8()
}

func fn10() error {
	return wrapperr.WithStack(fn7(), message1, message2, message3)
}

func fn11() error {
	sampleError := sampleErr{Data1: "data 1", Data2: "data 2"}
	return wrapperr.WithStackf(sampleError, "the cause is a special error of type %T", sampleError)
}

const (
	message1 = "message 1"
	message2 = "message 2"
	message3 = "message 3"
)
