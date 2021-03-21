package wrapperr_test

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/dc0d/wrapperr"
)

func show(err error) {
	terr, ok := err.(wrapperr.TracedErr)
	if !ok {
		fmt.Println(err)
		return
	}

	var stack wrapperr.Stack
	for _, v := range terr.Stack {
		if !strings.Contains(v.Loc.Func, "wrapperr") {
			v := v
			buildDefaultAnnotation(&v)
			stack = append(stack, v)
			break
		}
		v.Loc.File = ".../" + path.Base(v.Loc.File)
		stack = append(stack, v)
	}
	terr.Stack = stack

	fmt.Println(terr)
}

func showJSON(err error) {
	terr, ok := err.(wrapperr.TracedErr)
	if !ok {
		fmt.Println(err)
		return
	}

	var stack wrapperr.Stack
	for _, v := range terr.Stack {
		if !strings.Contains(v.Loc.Func, "wrapperr") {
			v := v
			buildDefaultAnnotation(&v)
			stack = append(stack, v)
			break
		}
		v.Loc.File = ".../" + path.Base(v.Loc.File)
		stack = append(stack, v)
	}
	terr.Stack = stack

	js, err := json.MarshalIndent(terr, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(js))
}

func buildDefaultAnnotation(note *wrapperr.Annotation) {
	const (
		rest = "..."
	)

	note.Loc.Line = 0
	note.Loc.File = rest
	note.Loc.Func = rest
	note.Message = "rest of stack"
}
