package wrapperr_test

import (
	"errors"

	"github.com/dc0d/wrapperr"
)

func end2() error {
	return wrapperr.WithStack(errors.New("OP-ERR"))
}

func caller2() error {
	return wrapperr.WithStack(end2(), "some notes on the call stack")
}

func begin2() error {
	return caller2()
}

func ExampleWithStack_withAnnotation() {
	err := begin2()
	show(err)

	// Output:
	// stack: .../example_2_test.go:10 github.com/dc0d/wrapperr_test.end2
	// >> .../example_2_test.go:14 github.com/dc0d/wrapperr_test.caller2 - some notes on the call stack
	// >> .../example_2_test.go:18 github.com/dc0d/wrapperr_test.begin2
	// >> .../example_2_test.go:22 github.com/dc0d/wrapperr_test.ExampleWithStack_withAnnotation
	// >> ...:0 ... - rest of stack
	// cause: OP-ERR
}
