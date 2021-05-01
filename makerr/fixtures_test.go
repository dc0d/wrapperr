package makerr_test

import (
	"time"

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

func checkScope(onEnter func(loc string), onExit func(elapsed time.Duration, loc string)) {
	defer makerr.TraceFn(onEnter)(onExit)
	time.Sleep(time.Millisecond * 50)
}
