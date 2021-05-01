package makerr

import "time"

func TraceFn(onEnter func(loc string)) func(onExit func(elapsed time.Duration, loc string)) {
	location := GetCaller(1)
	onEnter(location)

	startedAt := time.Now()
	return func(onExit func(elapsed time.Duration, loc string)) {
		onExit(time.Since(startedAt), location)
	}
}
