package wrapperr_test

import (
	"errors"

	"github.com/dc0d/wrapperr"
)

func end3() error {
	return wrapperr.WithStack(errors.New("OP-ERR"))
}

func caller3() error {
	return wrapperr.WithStack(end3(), "some notes on the call stack")
}

func begin3() error {
	return caller3()
}

func ExampleWithStack_toJSON() {
	err := begin3()
	showJSON(err)

	// Output:
	// {
	//   "stack": [
	//     ".../example_3_test.go:10 github.com/dc0d/wrapperr_test.end3",
	//     ".../example_3_test.go:14 github.com/dc0d/wrapperr_test.caller3 - some notes on the call stack",
	//     ".../example_3_test.go:18 github.com/dc0d/wrapperr_test.begin3",
	//     ".../example_3_test.go:22 github.com/dc0d/wrapperr_test.ExampleWithStack_toJSON",
	//     "...:0 ... - rest of stack"
	//   ],
	//   "cause": "OP-ERR"
	// }
}
