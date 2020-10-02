package here_test

import (
	"github.com/dc0d/here"
)

func firstFn() error {
	return here.Mark(rootCause)
}

func anotherFn() error {
	return here.Mark(rootCause)
}

func secondFn() error {
	return here.Mark(firstFn())
}

func thirdFn() error {
	return here.Mark(secondFn())
}

func callAnonymousFunc() error {
	fn := func() error { return here.Mark(rootCause) }
	return here.Mark(fn())
}
