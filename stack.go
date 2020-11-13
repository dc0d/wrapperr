package wrapperr

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type Stack []Annotation

func (stack Stack) String() string {
	var lines []string
	for _, note := range stack {
		lines = append(lines, note.String())
	}
	return strings.Join(lines, "\n>> ")
}

func WithStack(err error, message ...string) error {
	switch x := err.(type) {
	case TracedErr:
		if len(message) > 0 {
			calls := mark(3)

			for i, note := range x.Stack {
				if note.Loc == calls[0] {
					note.Message = message[0]
					x.Stack[i] = note
					break
				}
			}
		}
		return x
	default:
	}

	calls := mark(3)

	stack := make(Stack, len(calls))
	for i, call := range calls {
		stack[i] = Annotation{
			Loc: call,
		}
	}

	if len(message) > 0 {
		stack[0].Message = message[0]
	}

	return TracedErr{
		Stack: stack,
		Cause: err,
	}
}

type TracedErr struct {
	Stack Stack `json:"stack,omitempty"`
	Cause error `json:"cause,omitempty"`
}

func (terr TracedErr) Error() string {
	return "stack: " + terr.Stack.String() + "\ncause: " + fmt.Sprint(terr.Cause)
}

func (terr TracedErr) Unwrap() error { return terr.Cause }

func (terr TracedErr) MarshalJSON() ([]byte, error) {
	var payload struct {
		Stack Stack       `json:"stack,omitempty"`
		Cause interface{} `json:"cause,omitempty"`
	}

	payload.Stack = terr.Stack

	cause := terr.Unwrap()
	if cause != nil {
		js, err := json.Marshal(cause)

		switch {
		case err != nil:
			payload.Cause = cause.Error()
		case string(js) == "{}":
			payload.Cause = cause.Error()
		default:
			payload.Cause = cause
		}
	}

	return json.Marshal(payload)
}

func mark(skip int) (result []Loc) {
	const max = 64
	var pfuncs [max]uintptr
	n := runtime.Callers(skip, pfuncs[:])

	var calls []uintptr
	if n < 64 {
		calls = pfuncs[0:n]
	} else {
		calls = pfuncs[:]
	}

	frames := runtime.CallersFrames(calls)
	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		var funcName string
		fn := runtime.FuncForPC(frame.PC)
		if fn == nil {
			funcName = NotAvailableFuncName
		} else {
			funcName = fn.Name()
		}

		loc := Loc{
			File: frame.File,
			Line: frame.Line,
			Func: funcName,
		}

		result = append(result, loc)
	}

	return
}

const NotAvailableFuncName = "NOT_AVAILABLE"
