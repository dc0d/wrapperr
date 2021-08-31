package wrapperr

import (
	"runtime"
)

// GetCaller provides information about the caller - file, line number and the function/method.
func GetCaller(skip int) Loc {
	return locOf(skip + 1)
}

func locOf(skip int) Loc {
	if skip < 0 {
		skip = 0
	}
	skip++

	pc, file, line, ok := runtime.Caller(skip)

	name := NotAvailableFuncName
	f := runtime.FuncForPC(pc)
	if f != nil {
		name = f.Name()
	}

	var result Loc
	if ok {
		result.File = shortFilePath(file)
		result.Line = line
		result.Func = name
	}

	return result
}
