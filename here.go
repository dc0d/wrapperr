package here

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type Loc struct {
	Line int    `json:"line,omitempty"`
	File string `json:"file,omitempty"`
	Func string `json:"func,omitempty"`
}

func (loc Loc) String() string { return fmt.Sprintf(loc.File+":%d "+loc.Func, loc.Line) }

func (loc Loc) MarshalJSON() ([]byte, error) {
	var payload struct {
		File string `json:"file,omitempty"`
		Func string `json:"func,omitempty"`
	}
	payload.File = fmt.Sprintf(loc.File+":%d", loc.Line)
	payload.Func = loc.Func
	return json.Marshal(payload)
}

type Calls []Loc

func (calls Calls) String() string {
	var lines []string
	for _, loc := range calls {
		lines = append(lines, loc.String())
	}
	return strings.Join(lines, " >\n")
}

func Mark() (result Calls) {
	const max = 64
	var pfuncs [max]uintptr
	n := runtime.Callers(2, pfuncs[:])
	calls := pfuncs[0:n]

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
