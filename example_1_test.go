package wrapperr_test

import (
	"errors"

	"github.com/dc0d/wrapperr"
)

func end1() error {
	return wrapperr.WithStack(errors.New("OP-ERR"))
}

func caller1() error {
	return end1()
}

func begin1() error {
	return caller1()
}

func ExampleWithStack() {
	err := begin1()
	show(err)

	// Output:
	// stack: .../example_1_test.go:10 github.com/dc0d/wrapperr_test.end1
	// >> .../example_1_test.go:14 github.com/dc0d/wrapperr_test.caller1
	// >> .../example_1_test.go:18 github.com/dc0d/wrapperr_test.begin1
	// >> .../example_1_test.go:22 github.com/dc0d/wrapperr_test.ExampleWithStack
	// >> ...:0 ... - rest of stack
	// cause: OP-ERR
}
