package here

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strings"
)

func Mark(err error) error {
	lc := Here(WithSkip(2), WithShortNames())

	loc := cloc(lc)

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

// Here returns the location of a call, inside a caller function or a caller of the caller function and so on.
// The minimum number of skips is 1 (the direct caller of Here function).
func Here(options ...hereOption) (result Loc) {
	fnName := NotAvailableFuncName
	result.Func = fnName

	opt := newHereOptions()
	for _, op := range options {
		if op == nil {
			continue
		}
		opt = op(opt)
	}

	pc, file, line, ok := runtime.Caller(opt.skip)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fnName = fn.Name()

		if opt.shortFunc {
			fnName = path.Base(fnName)
		}
	}

	if opt.shortFile {
		file = path.Base(file)
	}

	result.File = file
	result.Line = line
	result.Func = fnName

	return
}

func WithSkip(skip int) hereOption {
	return func(opt hereOptions) hereOptions {
		if skip < 1 {
			return opt
		}

		opt.skip = skip

		return opt
	}
}

func WithShortNames() hereOption {
	return func(opt hereOptions) hereOptions {
		opt.shortFunc = true
		opt.shortFile = true

		return opt
	}
}

func WithShortFunc() hereOption {
	return func(opt hereOptions) hereOptions {
		opt.shortFunc = true

		return opt
	}
}

func WithShortFile() hereOption {
	return func(opt hereOptions) hereOptions {
		opt.shortFile = true

		return opt
	}
}

type hereOption func(hereOptions) hereOptions

type hereOptions struct {
	skip      int
	shortFunc bool
	shortFile bool
}

func newHereOptions() hereOptions {
	return hereOptions{
		skip: 1,
	}
}

type Loc cloc

const NotAvailableFuncName = "NOT_AVAILABLE"
