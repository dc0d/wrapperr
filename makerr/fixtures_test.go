package makerr_test

import (
	"github.com/dc0d/wrapperr/makerr"
)

func returnsExactLocation() (result string) {
	result = makerr.GetCaller(0)
	return
}

func toBeCalled() string {
	return makerr.GetCaller(1)
}

func returnsCallerLocation() string {
	return toBeCalled()
}
