package wrapperr_test

import "github.com/dc0d/wrapperr"

func sampleLoc() wrapperr.Loc {
	var loc wrapperr.Loc
	loc.File = "file"
	loc.Line = 10
	loc.Func = "fn"

	return loc
}
