package wrapperr_test

import (
	"github.com/dc0d/wrapperr"
)

func fn1() error {
	return wrapperr.WithStack(rootCause)
}

func fn2() error {
	return fn1()
}

func fn3() error {
	return wrapperr.WithStack(rootCause, "message 3")
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
