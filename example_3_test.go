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
	//     {
	//       "loc": {
	//         "file": ".../example_3_test.go:10",
	//         "func": "github.com/dc0d/wrapperr_test.end3"
	//       }
	//     },
	//     {
	//       "loc": {
	//         "file": ".../example_3_test.go:14",
	//         "func": "github.com/dc0d/wrapperr_test.caller3"
	//       },
	//       "message": "some notes on the call stack"
	//     },
	//     {
	//       "loc": {
	//         "file": ".../example_3_test.go:18",
	//         "func": "github.com/dc0d/wrapperr_test.begin3"
	//       }
	//     },
	//     {
	//       "loc": {
	//         "file": ".../example_3_test.go:22",
	//         "func": "github.com/dc0d/wrapperr_test.ExampleWithStack_toJSON"
	//       }
	//     },
	//     {
	//       "loc": {
	//         "file": "...:0",
	//         "func": "..."
	//       },
	//       "message": "rest of stack"
	//     }
	//   ],
	//   "cause": "OP-ERR"
	// }
}
