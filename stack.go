package wrapperr

import (
	"encoding/json"
	"fmt"
	"path"
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

// WithStack wraps the error inside another error. The original error is the root cause and
// can be unwrapped using the standard errors.Unwrap(...) function. The wrapper error contains
// the call stack in addition to the original error. When printed, it provides the call stack and
// the cause in text format. If it gets marshaled as JSON, it will include the call stack
// as a collection of calls. Each call information includes the code file path, line number
// and the function name and package and the annotation (if provided).
func WithStack(err error, messages ...string) error {
	var msg string
	if len(messages) > 0 {
		msg = strings.Join(messages, " ")
	}

	return makeErr(err, msg)
}

func WithStackf(err error, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)

	return makeErr(err, msg)
}

func makeErr(err error, msg string) error {
	switch x := err.(type) {
	case TracedErr:
		if msg != "" {
			calls := mark(defaultCallerStackSkip)

			for i, note := range x.Stack {
				if note.Loc == calls[0] {
					note.Message = msg
					x.Stack[i] = note
					break
				}
			}
		}
		return x
	default:
	}

	calls := mark(defaultCallerStackSkip)

	stack := make(Stack, len(calls))
	for i, call := range calls {
		stack[i] = Annotation{
			Loc: call,
		}
	}

	if msg != "" {
		stack[0].Message = msg
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

func mark(skip int) []Loc {
	var pfuncs [maxCallerDepth]uintptr
	n := runtime.Callers(skip, pfuncs[:])

	var calls []uintptr
	if n < maxCallerDepth {
		calls = pfuncs[0:n]
	} else {
		calls = pfuncs[:]
	}

	var result []Loc
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
			File: shortFilePath(frame.File),
			Line: frame.Line,
			Func: funcName,
		}

		result = append(result, loc)
	}

	return result
}

func shortFilePath(fp string) string {
	return path.Join(path.Base(path.Dir(fp)), path.Base(fp))
}

const (
	NotAvailableFuncName = "NOT_AVAILABLE"
)

const (
	maxCallerDepth         = 64
	defaultCallerStackSkip = 4
)
