package makerr

import (
	"path"
	"runtime"
)

func GetCaller(skip int) string {
	return locOf(skip + 1).String()
}

func locOf(skip int) loc {
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

	var result loc
	if ok {
		result.File = shortFilePath(file)
		result.Line = line
		result.Func = name
	}

	return result
}

func shortFilePath(fp string) string {
	return path.Join(path.Base(path.Dir(fp)), path.Base(fp))
}

const (
	NotAvailableFuncName = "NOT_AVAILABLE"
)
