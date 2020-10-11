package here

import (
	"encoding/json"
	"fmt"
	"path"
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

func (loc *Loc) ShortenFile() { loc.File = path.Base(loc.File) }
func (loc *Loc) ShortenFunc() { loc.Func = path.Base(loc.Func) }

type Calls []Loc

func (calls Calls) String() string {
	var lines []string
	for _, loc := range calls {
		lines = append(lines, loc.String())
	}
	return strings.Join(lines, " >\n")
}

func Mark(options ...hereOption) (result Calls) {
	var opt = newHereOptions()
	for _, f := range options {
		opt = f(opt)
	}

	const max = 64
	var pfuncs [max]uintptr
	n := runtime.Callers(opt.skip, pfuncs[:])
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
		if opt.shortenFiles {
			loc.ShortenFile()
		}
		if opt.shortenFuncs {
			loc.ShortenFunc()
		}
		result = append(result, loc)
	}

	return
}

const NotAvailableFuncName = "NOT_AVAILABLE"

func WithShortFiles() hereOption {
	return func(opt hereOptions) hereOptions {
		opt.shortenFiles = true
		return opt
	}
}

func WithShortFuncs() hereOption {
	return func(opt hereOptions) hereOptions {
		opt.shortenFuncs = true
		return opt
	}
}

func WithSkip(skip int) hereOption {
	return func(opt hereOptions) hereOptions {
		if skip < 2 {
			return opt
		}

		opt.skip = skip
		return opt
	}
}

type hereOption func(hereOptions) hereOptions

type hereOptions struct {
	shortenFiles bool
	shortenFuncs bool
	skip         int
}

func newHereOptions() (res hereOptions) {
	res.skip = 2
	return
}
