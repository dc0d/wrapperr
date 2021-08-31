package wrapperr_test

import "github.com/dc0d/wrapperr"

func sampleLoc() wrapperr.Loc {
	return wrapperr.Loc{
		File: "/path/to/package/file.go",
		Line: 9,
		Func: "github.com/user/module/package.(*Struct).method",
	}
}
