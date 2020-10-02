package here

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strings"
)

func Mark(err error) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}

	fnName := "N/A"
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fnName = fn.Name()
	}

	loc := cloc{
		Line: line,
		File: path.Base(file),
		Func: path.Base(fnName),
	}

	mkerr := MarkerError{
		Calls: []cloc{loc},
		Cause: err,
	}
	mkerr.flatten()

	return mkerr
}

type MarkerError struct {
	Calls calls
	Cause error
}

func (m MarkerError) Error() string {
	return fmt.Sprintf("%v\n%v", m.Calls, m.Cause)
}

func (m MarkerError) Unwrap() error {
	return m.Cause
}

func (m MarkerError) MarshalJSON() ([]byte, error) {
	var payload struct {
		Calls calls
		Cause interface{}
	}

	payload.Calls = m.Calls

	cause := m.Unwrap()
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

func (m *MarkerError) flatten() {
	down, ok := m.Cause.(MarkerError)
	if !ok {
		return
	}
	m.Calls = append(m.Calls, down.Calls...)
	m.Cause = down.Unwrap()
}

type calls []cloc

func (c calls) String() string {
	var calls []string
	for _, call := range c {
		calls = append(calls, call.String())
	}
	return strings.Join(calls, "\n")
}

type cloc struct {
	Line int
	File string
	Func string
}

func (loc cloc) String() string {
	return fmt.Sprintf("[%v:%v %v]", loc.File, loc.Line, loc.Func)
}

func (loc cloc) MarshalJSON() ([]byte, error) {
	return []byte(`"` + loc.String() + `"`), nil
}
