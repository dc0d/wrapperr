package wrapperr_test

import (
	"github.com/dc0d/wrapperr"
)

func returnsExactLocation() (result string) {
	result = wrapperr.GetCaller(0).String()
	return
}

func toBeCalled() string {
	return wrapperr.GetCaller(1).String()
}

func returnsCallerLocation() string {
	return toBeCalled()
}
